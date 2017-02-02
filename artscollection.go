package artscollection

import "github.com/OpenGLAMTools/artscollection/storage"

type Loader interface {
	Load() *Collection
}

type Adder interface {
	Add(storage.Storager) error
}

type Updater interface {
	Update(storage.Storager) error
}

type Remover interface {
	Remove(storage.Storager)
}

type Collection struct {
	Artworks    []storage.Storager
	Title       string
	Description string
	DataFields  []storage.Field
}

type Item interface {
	Filepath() string
	GetStorage() storage.Storager
}
