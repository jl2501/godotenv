package godotenv

import (
	"io"
	"os"
	"path/filepath"
)

// minimum set of methods needed for a File type
type Filer interface {
	io.Closer
	io.Reader
	WriteString(string) (int, error)
	Sync() error
	Stat() (os.FileInfo, error)
}

// minimum set of methods needed to pass a filesytem type
// at least can be afero or billy - hopefully unknown others as well
type Filesyser interface {
	Open(string) (Filer, error)
	Create(string) (Filer, error)
	MkdirAll(string, os.FileMode) error
	Walk(root string, walkFn filepath.WalkFunc) error
	WriteFile(string, []byte, os.FileMode) error
}
