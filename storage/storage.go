// Package storage defines the API over interfaces, which are needed to store
// data.
package storage

type Identifier interface {
	ID() (id string, ok bool)
}

type Setter interface {
	Set(fieldID string, value interface{}) error
}

type Getter interface {
	Get(fieldID string) (interface{}, error)
}
