package artscollection

import "errors"

type Loader interface {
	Load() *Collection
}

type Adder interface {
	Add(*Item) error
}

type Updater interface {
	Update(*Item) error
}

type Remover interface {
	Remove(*Item)
}

type Collection struct {
	Artworks    []Item
	Title       string
	Description string
	DataFields  []Field
}

type Item interface {
	Filepath() string
	GetStorage() *Storage
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
