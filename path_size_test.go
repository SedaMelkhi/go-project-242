package code

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatSize_PlainBytes(t *testing.T) {
	require.Equal(t, "0B", FormatSize(0, false))
	require.Equal(t, "1023B", FormatSize(1023, false))
	require.Equal(t, "1024B", FormatSize(1024, false))
	require.Equal(t, "1524B", FormatSize(1524, false))
	require.Equal(t, "1524000B", FormatSize(1524000, false))
	require.Equal(t, "1524000000B", FormatSize(1524000000, false))
	require.Equal(t, "1524000000000B", FormatSize(1524000000000, false))
}

func TestFormatSize_HumanReadable(t *testing.T) {
	require.Equal(t, "0B", FormatSize(0, true))
	require.Equal(t, "1023B", FormatSize(1023, true))
	require.Equal(t, "1.0KB", FormatSize(1024, true))
	require.Equal(t, "1.5KB", FormatSize(1524, true))
	require.Equal(t, "1.5MB", FormatSize(1524000, true))
	require.Equal(t, "1.4GB", FormatSize(1524000000, true))
	require.Equal(t, "1.4TB", FormatSize(1524000000000, true))
}

func TestGetSize_File_AllTrue_HumanFalse(t *testing.T) {
	path := filepath.Join("testdata", "dir", "file.txt")
	info, err := os.Lstat(path)
	require.NoError(t, err)

	got, err := GetPathSize(path, false, false, true)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(info.Size(), false))
	require.Equal(t, want, got)
}

func TestGetSize_File_AllTrue_HumanTrue(t *testing.T) {
	path := filepath.Join("testdata", "dir", "file.txt")
	info, err := os.Lstat(path)
	require.NoError(t, err)

	got, err := GetPathSize(path, false, true, true)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(info.Size(), true))
	require.Equal(t, want, got)
}

func TestGetSize_Dir_AllTrue_HumanFalse(t *testing.T) {
	dir := filepath.Join("testdata")

	i1, err := os.Lstat(filepath.Join(dir, "file.txt"))
	require.NoError(t, err)
	i2, err := os.Lstat(filepath.Join(dir, "file2.txt"))
	require.NoError(t, err)

	sum := i1.Size() + i2.Size()

	got, err := GetPathSize(dir, false, false, true)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(sum, false))
	require.Equal(t, want, got)
}

func TestGetSize_Dir_AllTrue_HumanTrue(t *testing.T) {
	dir := filepath.Join("testdata")

	i1, err := os.Lstat(filepath.Join(dir, "file.txt"))
	require.NoError(t, err)
	i2, err := os.Lstat(filepath.Join(dir, "file2.txt"))
	require.NoError(t, err)

	sum := i1.Size() + i2.Size()

	got, err := GetPathSize(dir, false, true, true)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(sum, true))
	require.Equal(t, want, got)
}

func TestGetSize_HiddenFile_AllFalse_Ignored(t *testing.T) {
	path := filepath.Join("testdata", "hidden_file", ".file.txt")

	got, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(0, false))
	require.Equal(t, want, got)
}

func TestGetSize_HiddenFile_AllTrue_Included(t *testing.T) {
	path := filepath.Join("testdata", "hidden_file", ".file.txt")
	info, err := os.Lstat(path)
	require.NoError(t, err)

	got, err := GetPathSize(path, false, false, true)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(info.Size(), false))
	require.Equal(t, want, got)
}

func TestGetSize_HiddenDir_AllFalse_Ignored_HumanFalse(t *testing.T) {
	dir := filepath.Join("testdata", ".dir")

	got, err := GetPathSize(dir, false, false, false)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(0, false))
	require.Equal(t, want, got)
}

func TestGetSize_HiddenDir_AllFalse_Ignored_HumanTrue(t *testing.T) {
	dir := filepath.Join("testdata", ".dir")

	got, err := GetPathSize(dir, false, true, false)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(0, true))
	require.Equal(t, want, got)
}

func TestGetSize_HiddenDir_AllTrue_Included_HumanFalse(t *testing.T) {
	dir := filepath.Join("testdata", ".dir")

	i1, err := os.Lstat(filepath.Join(dir, "file.txt"))
	require.NoError(t, err)
	sum := i1.Size()

	got, err := GetPathSize(dir, false, false, true)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(sum, false))
	require.Equal(t, want, got)
}

func TestGetSize_HiddenDir_AllTrue_Included_HumanTrue(t *testing.T) {
	dir := filepath.Join("testdata", ".dir")

	i1, err := os.Lstat(filepath.Join(dir, "file.txt"))
	require.NoError(t, err)
	sum := i1.Size()

	got, err := GetPathSize(dir, false, true, true)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(sum, true))
	require.Equal(t, want, got)
}

func TestGetSize_Dir_Recursive_IncludesNestedFiles(t *testing.T) {
	dir := "testdata"

	i1, err := os.Lstat(filepath.Join(dir, "file.txt"))
	require.NoError(t, err)
	i2, err := os.Lstat(filepath.Join(dir, "file2.txt"))
	require.NoError(t, err)
	i3, err := os.Lstat(filepath.Join(dir, "dir", "file.txt"))
	require.NoError(t, err)
	i4, err := os.Lstat(filepath.Join(dir, "hidden_file", "dir", "file.txt"))
	require.NoError(t, err)

	sum := i1.Size() + i2.Size() + i3.Size() + i4.Size()

	got, err := GetPathSize(dir, true, false, false)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(sum, false))
	require.Equal(t, want, got)
}

func TestGetSize_Dir_Recursive_AllFalse_IncludesNestedNonHiddenDirs(t *testing.T) {
	dir := filepath.Join("testdata")

	i1, err := os.Lstat(filepath.Join(dir, "file.txt"))
	require.NoError(t, err)
	i2, err := os.Lstat(filepath.Join(dir, "file2.txt"))
	require.NoError(t, err)

	i3, err := os.Lstat(filepath.Join(dir, "dir", "file.txt"))
	require.NoError(t, err)

	i4, err := os.Lstat(filepath.Join(dir, "hidden_file", "dir", "file.txt"))
	require.NoError(t, err)

	sum := i1.Size() + i2.Size() + i3.Size() + i4.Size()

	got, err := GetPathSize(dir, true, false, false)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(sum, false))
	require.Equal(t, want, got)
}

func TestGetSize_Dir_Recursive_AllTrue_IncludesHiddenFilesAndHiddenDirs(t *testing.T) {
	dir := filepath.Join("testdata")

	i1, err := os.Lstat(filepath.Join(dir, "file.txt"))
	require.NoError(t, err)
	i2, err := os.Lstat(filepath.Join(dir, "file2.txt"))
	require.NoError(t, err)

	i3, err := os.Lstat(filepath.Join(dir, "dir", "file.txt"))
	require.NoError(t, err)
	i4, err := os.Lstat(filepath.Join(dir, "hidden_file", "dir", "file.txt"))
	require.NoError(t, err)

	i5, err := os.Lstat(filepath.Join(dir, ".dir", "file.txt"))
	require.NoError(t, err)

	i6, err := os.Lstat(filepath.Join(dir, "hidden_file", ".file.txt"))
	require.NoError(t, err)

	sum := i1.Size() + i2.Size() + i3.Size() + i4.Size() + i5.Size() + i6.Size()

	got, err := GetPathSize(dir, true, false, true)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(sum, false))
	require.Equal(t, want, got)
}

func TestGetSize_Dir_Recursive_AllTrue_HumanTrue_FormatsButDoesNotChangeSum(t *testing.T) {
	dir := filepath.Join("testdata")

	i1, err := os.Lstat(filepath.Join(dir, "file.txt"))
	require.NoError(t, err)
	i2, err := os.Lstat(filepath.Join(dir, "file2.txt"))
	require.NoError(t, err)

	i3, err := os.Lstat(filepath.Join(dir, "dir", "file.txt"))
	require.NoError(t, err)
	i4, err := os.Lstat(filepath.Join(dir, "hidden_file", "dir", "file.txt"))
	require.NoError(t, err)

	i5, err := os.Lstat(filepath.Join(dir, ".dir", "file.txt"))
	require.NoError(t, err)
	i6, err := os.Lstat(filepath.Join(dir, "hidden_file", ".file.txt"))
	require.NoError(t, err)

	sum := i1.Size() + i2.Size() + i3.Size() + i4.Size() + i5.Size() + i6.Size()

	got, err := GetPathSize(dir, true, true, true)
	require.NoError(t, err)

	want := fmt.Sprintf("%s", FormatSize(sum, true))
	require.Equal(t, want, got)
}
