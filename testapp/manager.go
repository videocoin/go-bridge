package testapp

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/videocoin/go-bridge/client"
)

var (
	actions = [...]TransferType{
		Deposit,
		Withdraw,
	}
)

func NewManager(
	logger *logrus.Entry,
	period time.Duration,
	rng *rand.Rand,
	db *DB,
	keys []*ecdsa.PrivateKey,
	client *client.Client,
	tokenbank, coinbank common.Address,
) *Manager {
	return &Manager{
		logger:    logger,
		period:    period,
		rng:       rng,
		db:        db,
		keys:      keys,
		client:    client,
		tokenbank: tokenbank,
		coinbank:  coinbank,
	}
}

type Manager struct {
	logger              *logrus.Entry
	period              time.Duration
	rng                 *rand.Rand
	db                  *DB
	keys                []*ecdsa.PrivateKey
	client              *client.Client
	tokenbank, coinbank common.Address
}

func (m *Manager) Run(ctx context.Context) error {
	ticker := time.NewTicker(m.period)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := m.Step(ctx); err != nil {
				m.logger.Errorf("testapp failed: %v", err)
			}
		}
	}
}

func (m *Manager) Step(ctx context.Context) error {
	m.rng.Shuffle(len(m.keys), func(i int, j int) {
		m.keys[i], m.keys[j] = m.keys[j], m.keys[i]
	})
	var (
		// allow to submit only half of actions concurrently
		n           = m.rng.Intn(len(m.keys) / 2)
		group, gctx = errgroup.WithContext(ctx)
		amount      = big.NewInt(100) // 100 wei, in order not to drain a banks if something goes wrong
	)
	for i := 0; i < n; i++ {
		key := m.keys[i]
		action := m.rng.Intn(1)
		group.Go(func() error {
			var (
				info client.TransferInfo
				err  error
			)
			start := time.Now()
			if action == 0 {
				info, err = m.client.WaitDeposit(gctx, key, m.tokenbank, amount)
			} else {
				info, err = m.client.WaitWithdraw(gctx, key, m.coinbank, amount)
			}
			completed := time.Since(start)
			tr := Transfer{
				Type:        actions[action],
				LocalHash:   info.LocalTxHash,
				ForeignHash: info.ForeignTxHash,
				Latency:     completed,
			}
			if err != nil {
				tr.State = Failed
			} else {
				tr.State = Success
			}
			if err := m.db.SaveTransfer(&tr); err != nil {
				m.logger.Errorf("failed to save transfer %v", tr.LocalHash)
				return err
			}
			return nil
		})
	}
	return group.Wait()
}
