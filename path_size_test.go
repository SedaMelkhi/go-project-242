package code

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatSize(t *testing.T){
	assert.Equal(t, "0B",FormatSize(0, false))
	assert.Equal(t,"0B", FormatSize(0, true))
	assert.Equal(t, "1023B",FormatSize(1023, false))
	assert.Equal(t,"1023B", FormatSize(1023, true))
	assert.Equal(t, "1024B",FormatSize(1024, false))
	assert.Equal(t, "1.0KB", FormatSize(1024, true))
	assert.Equal(t, "1524B", FormatSize(1524, false))
	assert.Equal(t, "1.5KB", FormatSize(1524, true))
	assert.Equal(t, "1524000B", FormatSize(1524000, false))
	assert.Equal(t, "1.5MB", FormatSize(1524000, true))
	assert.Equal(t, "1524000000B", FormatSize(1524000000, false), )
	assert.Equal(t, "1.4GB", FormatSize(1524000000, true))
	assert.Equal(t, "1524000000000B", FormatSize(1524000000000, false), )
	assert.Equal(t, "1.4TB", FormatSize(1524000000000, true))
}

func TestGetPathSize_FileInDir(t *testing.T) {
	path := "./testdata"
	fileInfo, err := os.Lstat(filepath.Join(path, "file.txt"))
	require.NoError(t, err)
	size := fileInfo.Size()
	fileInfo, err = os.Lstat(filepath.Join(path, "file2.txt"))
	require.NoError(t, err)
	size += fileInfo.Size()
	expected := fmt.Sprintf("%dB\t%s", size, path)
	result, err := GetSize(path, false)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
	result, err = GetSize(path, true)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetPathSize_File(t *testing.T) {
	path := "./testdata/dir/file.txt"
	fileInfo, err := os.Lstat(path)
	require.NoError(t, err)
	expected := fmt.Sprintf("%dB\t%s", fileInfo.Size(), path)
	result, err := GetSize(path, false)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
	result, err = GetSize(path, true)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}
