package code

import (
	"testing"
	"fmt"
	"os"
	"path/filepath"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPathSize_FileInDir(t *testing.T){
	path := "./testdata"
	fileInfo, err := os.Lstat(filepath.Join(path, "file.txt"))
	require.NoError(t, err)
	size := fileInfo.Size()
	fileInfo, err = os.Lstat(filepath.Join(path, "file2.txt"))
	require.NoError(t, err)
	size += fileInfo.Size()
	expected := fmt.Sprintf("%d\t%s", size, path)
	result, err := GetSize(path)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetPathSize_File(t *testing.T){
	path := "./testdata/dir/file.txt"
	fileInfo, err := os.Lstat(path)
	require.NoError(t, err)
	expected := fmt.Sprintf("%d\t%s", fileInfo.Size(), path)
	result, err := GetSize(path)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}