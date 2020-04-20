package blocks

import (
	"context"
	"encoding/binary"
	"io/ioutil"
	"math/big"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func discard() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(ioutil.Discard)
	return log
}

func TestBlockRead(t *testing.T) {
	type testCase struct {
		desc   string
		block  uint64
		expect *big.Int
	}

	for _, tc := range []testCase{
		{
			desc:   "Empty",
			expect: big.NewInt(0),
		},
		{
			desc:   "Set",
			block:  10,
			expect: big.NewInt(10),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			tmpfile, err := ioutil.TempFile("", "test-block-resource-")
			require.NoError(t, err)
			defer os.Remove(tmpfile.Name())

			if tc.block != 0 {
				require.NoError(t, binary.Write(tmpfile, binary.BigEndian, tc.block))
			}
			resource, err := NewWriterBlockResource(logrus.NewEntry(discard()), tmpfile.Name())
			require.NoError(t, err)

			block, err := resource.Get(context.TODO())
			require.NoError(t, err)
			require.Equal(t, tc.expect.Int64(), block.Int64())
		})
	}
}

func TestBlockWrite(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test-block-resource-")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	r1, err := NewWriterBlockResource(logrus.NewEntry(discard()), tmpfile.Name())
	require.NoError(t, err)

	expect := 10
	for i := 1; i <= expect; i++ {
		require.NoError(t, r1.Set(context.TODO(), big.NewInt(int64(i))))
	}

	r2, err := NewWriterBlockResource(logrus.NewEntry(discard()), tmpfile.Name())
	require.NoError(t, err)

	block, err := r2.Get(context.TODO())
	require.NoError(t, err)
	require.Equal(t, expect, int(block.Int64()))
}
