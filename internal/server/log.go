package server

import "sync"

type Log struct {
	mu      sync.RWMutex
	records []Record
}

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	record.Offset = c.Size()
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) GetByOffset(offset uint64) (Record, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if offset >= c.Size() {
		return Record{}, OffsetError{offset}
	}
	return c.records[offset], nil
}

func (c *Log) Size() uint64 {
	return uint64(len(c.records))
}
