package dirs

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestReadDirFiles(t *testing.T) {
	req := require.New(t)
	ctx := context.Background()

	testFunc := func(added, expected int) func(t *testing.T) {
		return func(t *testing.T) {
			dir, err := os.MkdirTemp("", "initial")
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(dir)

			for i := 0; i < added; i++ {
				_, tErr := os.CreateTemp(dir, "file")
				if tErr != nil {
					return
				}
			}
			for i := 0; i < added; i++ {
				_, tErr := os.CreateTemp(dir, ".")
				if tErr != nil {
					return
				}
			}
			dirInfo := NewDirInfo(dir)
			dirErr := dirInfo.ReadDir(ctx)

			req.NoError(dirErr)
			req.Equal(dir, dirInfo.Name)
			req.Equal(expected, len(dirInfo.Files))
		}
	}

	t.Run("empty folder", testFunc(0, 0))
	t.Run("1 file folder", testFunc(1, 1))
	t.Run("5 file folder", testFunc(5, 5))
	t.Run("10 file folder", testFunc(10, 10))
}

func TestReadDirFilesError(t *testing.T) {
	req := require.New(t)
	ctx := context.Background()

	dirInfo := NewDirInfo("test")
	dirErr := dirInfo.ReadDir(ctx)

	req.Error(dirErr)

}
