package filesync

import (
	"context"
	"go-sync-dirs/internal/dirs"
	"go-sync-dirs/internal/files"
	"go-sync-dirs/internal/logging"
	"os"
	"path/filepath"
	"sync"
)

func SyncDirs(ctx context.Context, initial, target string) {
	var mu sync.RWMutex
	var wg sync.WaitGroup

	dirInitial := dirs.NewDirInfo(initial)
	err := dirInitial.ReadDir(ctx)
	if err != nil {
		return
	}

	dirTarget := dirs.NewDirInfo(target)
	err = dirTarget.ReadDir(ctx)
	if err != nil {
		return
	}

OUTERLOOP:
	for _, initialFile := range dirInitial.Files {
		for _, targetFile := range dirTarget.Files {
			if initialFile.IsSameFile(targetFile) {
				continue OUTERLOOP
			}
		}

		wg.Add(1)
		go copyFile(ctx, initialFile, target, &mu, &wg)
	}
	wg.Wait()
}

func copyFile(ctx context.Context, file *files.FileInfo, target string, mu *sync.RWMutex, wg *sync.WaitGroup) {
	logger := logging.LoggerFromContext(ctx, "logger")
	mu.RLock()
	defer mu.RUnlock()
	defer wg.Done()

	absPAth, err := filepath.Abs(target)
	if err != nil {
		logger.Errorf("%e", err)
		return
	}

	fileContent, err := os.ReadFile(file.FullPath)
	if err != nil {
		logger.Errorf("%e", err)
		return
	}
	err = os.WriteFile(filepath.Join(absPAth, file.Name), fileContent, file.Mode)
	if err != nil {
		logger.Errorf("%e", err)
		return
	}

	logger.Infof("Copied file %s sized %d to target folder %s", file.Name, file.Size, target)
}
