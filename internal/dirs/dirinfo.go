package dirs

import (
	"context"
	"crypto/sha512"
	"fmt"
	"go-sync-dirs/internal/files"
	"go-sync-dirs/internal/logging"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type DirReader interface {
	ReadDir(ctx context.Context) error
}

type DirInfo struct {
	Name  string
	Files []*files.FileInfo
}

func NewDirInfo(name string) *DirInfo {
	return &DirInfo{
		Name: name,
	}
}

func (dir *DirInfo) ReadDir(ctx context.Context) error {
	var mu sync.RWMutex
	var wg sync.WaitGroup
	logger := logging.LoggerFromContext(ctx, "logger")

	allFiles, err := os.ReadDir(dir.Name)
	if err != nil {
		logger.Errorf("Can't read directory %s: %v", dir.Name, err)
		return err
	}

	dirFiles := make([]*files.FileInfo, 0, len(allFiles))
	for _, fileEntry := range allFiles {
		wg.Add(1)
		go func(fileEntry os.DirEntry) {
			defer wg.Done()
			fileInfo, fileErr := GenFileInfo(dir, fileEntry, &mu)
			if fileErr != nil {
				logger.Errorf("Can't process file %s: %e", fileInfo.Name, fileErr)
				return
			}
			if fileInfo == nil {
				return
			}

			dirFiles = append(dirFiles, fileInfo)
		}(fileEntry)

	}

	wg.Wait()
	dir.Files = dirFiles
	return nil
}

func GenFileInfo(dir *DirInfo, file os.DirEntry, mu *sync.RWMutex) (*files.FileInfo, error) {
	mu.RLock()
	defer mu.RUnlock()

	if file.IsDir() || strings.HasPrefix(file.Name(), ".") {
		return nil, nil
	}
	fullPath := filepath.Join(dir.Name, file.Name())
	fileContent, contentErr := os.ReadFile(fullPath)
	if contentErr != nil {
		return nil, contentErr
	}
	fileData, fileErr := file.Info()
	if fileErr != nil {
		return nil, fileErr
	}

	fileInfo := files.NewFileInfo(
		fullPath,
		file.Name(),
		fmt.Sprintf("%x", sha512.Sum512(fileContent)),
		fileData.Size(),
		fileData.Mode(),
	)

	return fileInfo, nil
}
