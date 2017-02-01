package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Txt struct {
	Fields   []Field
	Strings  map[string]string
	Integers map[string]int
	Bools    map[string]bool
}

var ErrTypeNotSupported = errors.New("Type is not supported!")
var ErrFieldNotSupported = errors.New("Field is not supported!")
var ErrWrongType = errors.New("Input has the wrong type!")

func NewTxtStorage() *Txt {
	return &Txt{
		Strings:  make(map[string]string),
		Integers: make(map[string]int),
		Bools:    make(map[string]bool),
	}
}

func (txt *Txt) Set(fieldID string, value interface{}) error {
	fileDef, ok := txt.GetField(fieldID)
	switch {
	case !ok:
		return ErrFieldNotSupported
	}
	switch t := value.(type) {
	case string:
		if fileDef.Type != "string" {
			return ErrWrongType
		}
		txt.Strings[fieldID] = t
	case int:
		if fileDef.Type != "int" {
			return ErrWrongType
		}
		txt.Integers[fieldID] = t
	case bool:
		if fileDef.Type != "bool" {
			return ErrWrongType
		}
		txt.Bools[fieldID] = t
	default:
		return ErrTypeNotSupported
	}
	return nil
}

// Get returns a value from the storage
func (txt *Txt) Get(fieldID string) (interface{}, bool) {
	var out interface{}
	f, ok := txt.GetField(fieldID)
	t := f.Type
	if !ok {
		return out, false
	}
	switch t {
	case "string":
		out = txt.Strings[fieldID]
	case "int":
		out = txt.Integers[fieldID]
	case "bool":
		out = txt.Bools[fieldID]
	default:
		return out, false
	}
	return out, true
}

func (txt *Txt) GetString(fieldID string) (string, bool) {
	if !txt.checkType(fieldID, "string") {
		var out string
		return out, false
	}
	return txt.Strings[fieldID], true
}

func (txt *Txt) GetInt(fieldID string) (int, bool) {
	if !txt.checkType(fieldID, "int") {
		var out int
		return out, false
	}
	return txt.Integers[fieldID], true
}

func (txt *Txt) GetBool(fieldID string) (bool, bool) {
	if !txt.checkType(fieldID, "bool") {
		var out bool
		return out, false
	}
	return txt.Bools[fieldID], true
}

func (txt *Txt) checkType(fieldID, fieldType string) bool {
	f, ok := txt.GetField(fieldID)
	if !ok {
		return false
	}
	return f.Type == fieldType
}

// GetField returns the field
func (txt *Txt) GetField(fieldID string) (Field, bool) {
	for _, f := range txt.Fields {
		if f.ID() == fieldID {
			return f, true
		}
	}
	return Field{}, false
}

// AddField adds a field to the storage
func (txt *Txt) AddField(f Field) error {
	txt.Fields = append(txt.Fields, f)
	return nil
}

// Marshal is used for a json output
func (txt *Txt) Marshal() ([]byte, error) {
	return json.MarshalIndent(txt, "", "  ")
}

// Unmarshal uses json to set all the data into the Txt instance.
func (txt *Txt) Unmarshal(b []byte) error {
	return json.Unmarshal(b, txt)
}

// LoadTxt loads a Txt storage from a file
func LoadTxt(filename string) (*Txt, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	txt := NewTxtStorage()
	err = txt.Unmarshal(b)
	if err != nil {
		return nil, err
	}
	return txt, nil
}

// WriteTxt writes a Txt Storage to a file
func WriteTxt(txt *Txt, filename string) error {
	b, err := txt.Marshal()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0777)
}
