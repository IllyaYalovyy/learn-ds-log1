package data

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testMessages = []string{"Short message", "A very long test message", "", "The End"}
)

func TestStore_AppendRead(t *testing.T) {
	s := testCreateStore(t)
	testAppendData(t, s)
	testRead(t, s)
}

func TestStore_ReadAt(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	s := testCreateStore(t)
	testAppendData(t, s)

	message := testMessages[0]
	data := make([]byte, len(message))
	numBytes, err := s.ReadAt(data, recordLenSize)
	require.NoError(err)
	assert.Equal(message, string(data), "written data")
	assert.Equal(len(message), numBytes)
}

func testCreateStore(t *testing.T) *Store {
	t.Helper()
	require := require.New(t)

	testDir := t.TempDir()
	f, err := os.Create(filepath.Join(testDir, "log_store.bin"))
	require.NoError(err)
	s, err := NewStore(f)
	require.NoError(err)
	return s
}

func testAppendData(t *testing.T, s *Store) {
	t.Helper()
	require := require.New(t)
	assert := assert.New(t)

	var expectedPos int64
	for _, message := range testMessages {
		n, pos, err := s.Append([]byte(message))
		require.NoError(err)
		assert.Equal(expectedPos, pos, "record position")
		expectedRecordSize := int64(len(message) + recordLenSize)
		assert.Equal(expectedRecordSize, n, "record size")
		expectedPos += expectedRecordSize
	}
}

func testRead(t *testing.T, s *Store) {
	t.Helper()
	require := require.New(t)
	assert := assert.New(t)

	var readPos int64
	for _, message := range testMessages {
		data, err := s.Read(readPos)
		require.NoError(err)
		assert.Equal(message, string(data), "written data")
		readPos += int64(len(data) + recordLenSize)
	}
}

func TestStore_CloseOpenStore(t *testing.T) {
	require := require.New(t)
	s := testCreateStore(t)
	testAppendData(t, s)

	fileName := s.Name()

	err := s.Close()
	require.NoError(err)

	f, err := os.Open(fileName)
	require.NoError(err)

	s, err = NewStore(f)
	require.NoError(err)
	testRead(t, s)
}
