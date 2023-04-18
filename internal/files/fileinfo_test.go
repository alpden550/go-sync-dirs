package files

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileCompare(t *testing.T) {
	req := require.New(t)

	cases := map[string]struct {
		file1    *FileInfo
		file2    *FileInfo
		expected bool
	}{
		"compare two equal": {
			NewFileInfo("path", "name", "sha512", 1, 0644),
			NewFileInfo("new/path", "name", "sha512", 1, 0644),
			true,
		},
		"compare two diff name": {
			NewFileInfo("path", "name", "sha512", 1, 0644),
			NewFileInfo("new/path", "new name", "sha512", 1, 0644),
			false,
		},
		"compare two diff sha512": {
			NewFileInfo("path", "name", "sha512", 1, 0644),
			NewFileInfo("new/path", "name", "sha256", 1, 0644),
			false,
		},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			file1, file2 := testCase.file1, testCase.file2
			res := file1.IsSameFile(file2)
			req.Equal(testCase.expected, res)
		})
	}
}
