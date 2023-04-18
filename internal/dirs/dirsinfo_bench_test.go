package dirs

import (
	"context"
	"log"
	"os"
	"testing"
)

func BenchmarkReadDirFiles(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		dir, err := os.MkdirTemp("", "initial")
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < b.N; i++ {
			_, tErr := os.CreateTemp(dir, "file")
			if tErr != nil {
				return
			}
		}

		dirInfo := NewDirInfo(dir)
		dirErr := dirInfo.ReadDir(ctx)
		if dirErr != nil {
			return
		}

		err = os.RemoveAll(dir)
		if err != nil {
			return
		}
	}
}
