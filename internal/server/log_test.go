package server_test

import (
	"testing"

	"github.com/IllyaYalovyy/learn-ds-log1/internal/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppendGet(t *testing.T) {
	logServer := server.New()
	testRecord := server.Record{Value: []byte{'h', 'a', 'l', 'l', 'o'}}
	testOffset, err := logServer.Append(testRecord)
	require.NoError(t, err)
	actualRecord, err := logServer.GetByOffset(testOffset)
	require.NoError(t, err)
	assert.Equal(t, testRecord, actualRecord)
}

func TestInvalidOffset(t *testing.T) {
	logServer := server.New()
	_, err := logServer.GetByOffset(123)
	assert.EqualError(t, err, "invalid offset:123")
}
