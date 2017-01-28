// Package storage defines the API over interfaces, which are needed to store
// data.
package storage

// Field abstracts the definition of a field. The supportet Properties
// are provided by the implementation.
type Field interface {
	ID() string
}

// Setter sets the data of a field
type Setter interface {
	Set(fieldID string, value interface{}) error
}

// Getter is used to get data out of the storage
type Getter interface {
	Get(fieldID string) (interface{}, bool)
	GetInt(fieldID string) (int, bool)
	GetString(fieldID string) (string, bool)
}

// Storager combines the interfaces
type Storager interface {
	Setter
	Getter
	AddField(Field) error
	GetFields() ([]Field, error)
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}
