package filesync

import (
	"context"
	"log"
	"os"
	"testing"
)

func BenchmarkSyncDirs(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		dirInitial, err := os.MkdirTemp("", "initial")
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < b.N; i++ {
			_, tErr := os.CreateTemp(dirInitial, "file")
			if tErr != nil {
				return
			}
		}

		dirTarget, err := os.MkdirTemp("", "target")
		if err != nil {
			log.Fatal(err)
		}

		SyncDirs(ctx, dirInitial, dirTarget)

		_ = os.RemoveAll(dirInitial)
		_ = os.RemoveAll(dirTarget)
	}
}
