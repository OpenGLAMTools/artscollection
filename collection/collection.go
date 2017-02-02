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
	fpath    string
}

// NewCollection returns an empty collection
func NewCollection(fpath string) *Collection {
	return &Collection{
		Storages: make(map[string]storage.Storager),
		fpath:    fpath,
	}
}

// GetItem returns a storage from the collection
func (c *Collection) GetItem(ID string) (storage.Storager, bool) {
	i, ok := c.Storages[ID]
	return i, ok
}

func (c *Collection) WriteItem(ID string, s storage.Storager) error {
	itemPath := filepath.Join(
		c.Path(),
		ID,
		StorageFile,
	)
	return storage.Write(s, itemPath)
}

// Marshal returns the whole collection as json object
func (c *Collection) Marshal() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

// Path returns the root path to the collection
func (c *Collection) Path() string {
	return c.fpath
}

func Load(fpath string) (*Collection, error) {
	c := NewCollection(fpath)
	confFile := filepath.Join(
		fpath,
		FieldsConfFile,
	)
	var err error
	c.Fields, err = storage.LoadFields(confFile)
	if err != nil {
		return c, err
	}
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
		if len(s.Fields) == 0 {
			s.Fields = c.Fields
		}
		c.Storages[fi.Name()] = s
	}

	return c, nil
}
