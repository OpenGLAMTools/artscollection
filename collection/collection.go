package collection

import (
	"encoding/json"

	"github.com/OpenGLAMTools/artscollection/storage"
)

// Collection defines the fields for a storage that fields are used as
// default for new storages.
type Collection struct {
	Fields   []*storage.Field
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
	return json.Marshal(c)
}
