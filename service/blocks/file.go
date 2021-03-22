package blocks

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
	"math/big"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/videocoin/go-bridge/service"
)

var (
	ErrIncompleteWrite = errors.New("incomplete write")

	_ service.LastBlockAccess = (*WriterBlockResource)(nil)
)

const (
	bufsize = 8
)

func NewWriterBlockResource(log *logrus.Entry, filename string) (*WriterBlockResource, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	resource := &WriterBlockResource{
		log:      log,
		filename: filename,
		writer:   f,
		block:    12071765,
	}
	/*
		if err := binary.Read(f, binary.BigEndian, &resource.block); err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}
	*/

	return resource, nil
}

type WriterBlockResource struct {
	log      *logrus.Entry
	filename string
	block    uint64
	writer   io.WriterAt
}

func (w *WriterBlockResource) Get(context.Context) (*big.Int, error) {
	return new(big.Int).SetUint64(w.block), nil
}

func (w *WriterBlockResource) Set(_ context.Context, number *big.Int) error {
	w.block = number.Uint64()

	buf := make([]byte, bufsize)
	binary.BigEndian.PutUint64(buf, w.block)
	n, err := w.writer.WriteAt(buf, 0)
	if err != nil {
		return err
	}
	if n != bufsize {
		return ErrIncompleteWrite
	}
	return nil
}
