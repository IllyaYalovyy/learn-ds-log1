package data

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	bigEndian = binary.BigEndian
)

const recordLenSize = 8

type Store struct {
	*os.File
	mu   sync.Mutex
	buf  *bufio.Writer
	size int64
}

func NewStore(file *os.File) (*Store, error) {
	fileInfo, err := os.Stat(file.Name())
	if err != nil {
		return nil, err
	}
	size := fileInfo.Size()
	buf := bufio.NewWriter(file)
	return &Store{File: file, size: size, buf: buf}, nil
}

func (s *Store) Append(p []byte) (n int64, pos int64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	pos = s.size
	err = binary.Write(s.buf, bigEndian, uint64(len(p)))
	if err != nil {
		return
	}
	writtenBytes, err := s.buf.Write(p)
	if err != nil {
		return
	}
	totalRecordSize := int64(recordLenSize + writtenBytes)
	s.size += totalRecordSize
	return totalRecordSize, pos, nil
}

func (s *Store) Read(pos int64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.buf.Flush()
	if err != nil {
		return nil, err
	}
	sizeBinary := make([]byte, recordLenSize)
	_, err = s.File.ReadAt(sizeBinary, pos)
	if err != nil {
		return nil, err
	}
	recordDataSize := bigEndian.Uint64(sizeBinary)
	data := make([]byte, recordDataSize)
	_, err = s.File.ReadAt(data, pos+recordLenSize)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Store) ReadAt(p []byte, off int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.buf.Flush()
	if err != nil {
		return 0, err
	}
	return s.File.ReadAt(p, off)
}

func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.buf.Flush()
	if err != nil {
		return err
	}
	return s.File.Close()
}
