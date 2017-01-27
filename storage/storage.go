// Package storage defines the API over interfaces, which are needed to store
// data.
package storage

// Field abstracts the definition of a field. The supportet Properties
// are provided by the implementation.
type Field interface {
	ID() string
}

// Properties abstracts a slice of Field.
type Properties interface {
	AddField(Field) error
	GetFields() ([]Field, error)
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

// Marshaler is used for the external API. It marshals the data of the
// storage.
type Marshaler interface {
	Marshal() ([]byte, error)
}

// Unmarshaler is used for the external API. It unmarshals the input into
// the storage.
type Unmarshaler interface {
	Unmarshal([]byte) error
}

// Storager combines the interfaces
type Storager interface {
	Properties
	Setter
	Getter
	Marshaler
	Unmarshaler
}
