package godotenv

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

type AferoFs struct {
	Afs afero.Fs
}

func (a AferoFs) Create(name string) (Filer, error) {
	return a.Afs.Create(name)
}

func NewOsFs() Filesyser {
	return AferoFs{Afs: afero.NewOsFs()}
}

func NewMemMapFs() Filesyser {
	return AferoFs{Afs: afero.NewMemMapFs()}
}

func (a AferoFs) Open(filename string) (Filer, error) {
	return a.Afs.Open(filename)
}

func (a AferoFs) Walk(root string, walkFn filepath.WalkFunc) error {
	return afero.Walk(a.Afs, root, walkFn)
}

func (a AferoFs) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return afero.WriteFile(a.Afs, filename, data, perm)
}

func (a AferoFs) MkdirAll(path string, perm os.FileMode) error {
	return a.Afs.MkdirAll(path, perm)
}
