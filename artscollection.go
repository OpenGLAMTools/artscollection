package artscollection

import (
	"errors"

	"github.com/OpenGLAMTools/artscollection/storage"
)

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
	DataFields  []Field
}

type Item interface {
	Filepath() string
	GetStorage() storage.Storager
}

// Storage defines all the supported types and represents the API
type Storage struct {
	Fields   []Field
	Strings  map[string]string
	Integers map[string]int
	Bools    map[string]bool
}

// Field defines all fields for an item.
type Field struct {
	// ID is a unique identifier for the field
	ID string
	// Name of the field, does not have to be unique
	Name string
	// Type of the field must match with the defined types
	// inside of the Data type.
	Type string
	// The name of a group the field is used.
	Group string
	// Items are ordered inside the group.
	Order int
	// If not nil that values are used for a dropdown list
	Select *[]string
}

var ErrTypeNotSupported = errors.New("Type is not supported!")

func (d *Storage) Set(fieldID string, value interface{}) error {
	switch t := value.(type) {
	case string:
		d.Strings[fieldID] = t
	case int:
		d.Integers[fieldID] = t
	case bool:
		d.Bools[fieldID] = t
	default:
		return ErrTypeNotSupported
	}
	return nil
}
