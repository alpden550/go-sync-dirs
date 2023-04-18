package filesync

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestSyncDirs(t *testing.T) {
	req := require.New(t)
	ctx := context.Background()

	testFunc := func(added, expected int) func(t *testing.T) {
		return func(t *testing.T) {
			dirInitial, err := os.MkdirTemp("", "initial")
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(dirInitial)

			for i := 0; i < added; i++ {
				_, tErr := os.CreateTemp(dirInitial, "file")
				if tErr != nil {
					return
				}
			}

			dirTarget, err := os.MkdirTemp("", "target")
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(dirTarget)

			SyncDirs(ctx, dirInitial, dirTarget)
			targetFiles, _ := os.ReadDir(dirTarget)

			req.Equal(expected, len(targetFiles))
		}
	}

	t.Run("empty initial", testFunc(0, 0))
	t.Run("1 initial", testFunc(1, 1))
	t.Run("3 initial", testFunc(3, 3))
	t.Run("10 initial", testFunc(10, 10))
}
