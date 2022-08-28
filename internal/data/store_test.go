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
	require := require.New(t)
	assert := assert.New(t)
	testDir := t.TempDir()
	f, err := os.Create(filepath.Join(testDir, "log_store.bin"))
	require.NoError(err)
	s, err := NewStore(f)
	require.NoError(err)

	var expectedPos int64
	for _, message := range testMessages {

		n, pos, err := s.Append([]byte(message))
		require.NoError(err)
		assert.Equal(expectedPos, pos, "record position")
		expectedRecordSize := int64(len(message) + recordLenSize)
		assert.Equal(expectedRecordSize, n, "record size")
		expectedPos += expectedRecordSize
	}

	var readPos int64
	for _, message := range testMessages {
		data, err := s.Read(readPos)
		require.NoError(err)
		assert.Equal(message, string(data), "written data")
		readPos += int64(len(data) + recordLenSize)
	}
}
