// Package storage defines the API over interfaces, which are needed to store
// data.
package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Setter sets the data of a field
type Setter interface {
	Set(fieldID string, value interface{}) error
}

// Getter is used to get data out of the storage
type Getter interface {
	Get(fieldID string) (interface{}, bool)
	GetInt(fieldID string) (int, bool)
	GetString(fieldID string) (string, bool)
	GetBool(fieldID string) (bool, bool)
	//GetTaxonomyTerms(taxonomy string) []string
}

// Storager combines the interfaces
type Storager interface {
	Setter
	Getter
	//AddField(Field) error
	GetFields() (Fields, error)
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

// Load loads data from a file into a Storager
func Load(filename string, s Storager) error {
	// Check if file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = s.Unmarshal(b)
	if err != nil {
		return err
	}
	return nil
}

// Write writes a Storager to a file
func Write(txt Storager, filename string) error {
	// ensure dir
	os.MkdirAll(filepath.Dir(filename), 0777)
	b, err := txt.Marshal()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0777)
}
