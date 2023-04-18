package files

import "io/fs"

type FileInfo struct {
	FullPath string
	Name     string
	Size     int64
	Sha512   string
	Mode     fs.FileMode
}

func NewFileInfo(path, name, sha512 string, size int64, mode fs.FileMode) *FileInfo {
	return &FileInfo{
		FullPath: path,
		Name:     name,
		Sha512:   sha512,
		Size:     size,
		Mode:     mode,
	}
}

func (f *FileInfo) IsSameFile(c *FileInfo) bool {
	return f.Name == c.Name && f.Sha512 == c.Sha512
}
