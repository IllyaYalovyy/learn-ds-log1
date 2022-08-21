package server

import "fmt"

type OffsetError struct {
	offset uint64
}

func (err OffsetError) Error() string {
	return fmt.Sprint("invalid offset:", err.offset)
}
