package collection

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/OpenGLAMTools/artscollection/storage"
)

// StorageFile is the default name
var StorageFile = "data.json"

// FieldsConfFile defines the default values for the storages
var FieldsConfFile = "conf.yaml"

// Collection defines the fields for a storage that fields are used as
// default for new storages.
type Collection struct {
	Fields   storage.Fields
	Storages map[string]storage.Storager
}

// NewCollection returns an empty collection
func NewCollection() *Collection {
	return &Collection{
		Storages: make(map[string]storage.Storager),
	}
}

// GetItem returns a storage from the collection
func (c *Collection) GetItem(ID string) (storage.Storager, bool) {
	i, ok := c.Storages[ID]
	return i, ok
}

// Marshal returns the whole collection as json object
func (c *Collection) Marshal() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

func Load(fpath string) (*Collection, error) {
	c := NewCollection()
	dir, err := ioutil.ReadDir(fpath)
	if err != nil {
		return c, err
	}
	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		storagePath := filepath.Join(
			fpath,
			fi.Name(),
			StorageFile,
		)
		s := storage.NewTxtStorage()
		err = storage.Load(storagePath, s)
		if err != nil {
			return c, err
		}
		c.Storages[fi.Name()] = s
	}
	confFile := filepath.Join(
		fpath,
		FieldsConfFile,
	)
	c.Fields, err = storage.LoadFields(confFile)
	if err != nil {
		return c, err
	}
	return c, nil
}
