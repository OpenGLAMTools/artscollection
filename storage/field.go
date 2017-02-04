package storage

import (
	"io/ioutil"
	"os"
	"sort"

	yaml "gopkg.in/yaml.v2"
)

// Field defines all fields for an item.
type Field struct {
	// Key is a unique identifier for the field
	Key string `yaml:"key"`
	// Name of the field, does not have to be unique
	Name string `yaml:"name"`
	// Type of the field must match with the defined types
	// inside of the Data type.
	Type string `yaml:"type"`
	// The name of a group the field is used.
	Group string `yaml:"group"`
	// Items are ordered inside the group.
	Order int `yaml:"order"`
	// Index specifies if that field should be indexed.
	Index bool `yaml:"index"`
	// If not nil that values are used for a dropdown list
	//Select []string
}

func (f *Field) ID() string {
	return f.Key
}

type Fields []Field

func (f Fields) Len() int           { return len(f) }
func (f Fields) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f Fields) Less(i, j int) bool { return f[i].Order < f[j].Order }

func LoadFields(fpath string) (Fields, error) {
	var fields []Field
	// Check if file exists
	_, err := os.Stat(fpath)
	if err == nil {
		confb, err := ioutil.ReadFile(fpath)
		if err != nil {
			return fields, err
		}

		err = yaml.Unmarshal(confb, &fields)
		if err != nil {
			return fields, err
		}
	}
	sort.Sort(Fields(fields))
	return fields, nil
}
