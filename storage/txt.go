package storage

import (
	"encoding/json"
	"errors"
)

// Txt is a txt based storager
type Txt struct {
	Fields       []Field             `json:"fields"`
	Strings      map[string]string   `json:"strings"`
	Integers     map[string]int      `json:"integers"`
	Bools        map[string]bool     `json:"bools,omitempty"`
	SliceStrings map[string][]string `json:"sclicestrings"`
}

const (
	TypeString      = "string"
	TypeInt         = "int"
	TypeBool        = "bool"
	TypeSliceString = "SliceString"
)

// ErrTypeNotSupported is used when the type of the interface
// is not able to stored inside the txt.
var ErrTypeNotSupported = errors.New("Type is not supported!")

// ErrFieldNotSupported is used, when the field is not inside the txt
// storage.
var ErrFieldNotSupported = errors.New("Field is not supported!")

// ErrWrongType is used when the given type does not match to the fieldID.
var ErrWrongType = errors.New("Input has the wrong type!")

// NewTxtStorage creates an empty Txt storage.
func NewTxtStorage() *Txt {
	return &Txt{
		Strings:      make(map[string]string),
		Integers:     make(map[string]int),
		Bools:        make(map[string]bool),
		SliceStrings: make(map[string][]string),
	}
}

// Set the value of a field
func (txt *Txt) Set(fieldID string, value interface{}) error {
	fileDef, ok := txt.GetField(fieldID)
	switch {
	case !ok:
		return ErrFieldNotSupported
	}
	switch t := value.(type) {
	case string:
		if fileDef.Type != TypeString {
			return ErrWrongType
		}
		txt.Strings[fieldID] = t
	case int:
		if fileDef.Type != TypeInt {
			return ErrWrongType
		}
		txt.Integers[fieldID] = t
	case bool:
		if fileDef.Type != TypeBool {
			return ErrWrongType
		}
		txt.Bools[fieldID] = t
	case []string:
		if fileDef.Type != TypeSliceString {
			return ErrWrongType
		}
		txt.SliceStrings[fieldID] = t
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
	case TypeString:
		out = txt.Strings[fieldID]
	case TypeInt:
		out = txt.Integers[fieldID]
	case TypeBool:
		out = txt.Bools[fieldID]
	case TypeSliceString:
		out = txt.SliceStrings[fieldID]
	default:
		return out, false
	}
	return out, true
}

// GetString returns the field value if it is a string
func (txt *Txt) GetString(fieldID string) (string, bool) {
	if !txt.checkType(fieldID, TypeString) {
		var out string
		return out, false
	}
	return txt.Strings[fieldID], true
}

func (txt *Txt) GetInt(fieldID string) (int, bool) {
	if !txt.checkType(fieldID, TypeInt) {
		var out int
		return out, false
	}
	return txt.Integers[fieldID], true
}

func (txt *Txt) GetBool(fieldID string) (bool, bool) {
	if !txt.checkType(fieldID, TypeBool) {
		var out bool
		return out, false
	}
	return txt.Bools[fieldID], true
}

func (txt *Txt) GetSliceString(fieldID string) ([]string, bool) {
	if !txt.checkType(fieldID, TypeSliceString) {
		var out []string
		return out, false
	}
	return txt.SliceStrings[fieldID], true
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

// GetFields returns all fields
func (txt *Txt) GetFields() (Fields, error) {
	return txt.Fields, nil
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
	err := json.Unmarshal(b, txt)
	if err != nil {
		return err
	}
	txt.Clean()
	return nil
}

// Clean removes all values which are not defined as field.
func (txt *Txt) Clean() {
	for key := range txt.Strings {
		f, ok := txt.GetField(key)
		if ok == false || f.Type != TypeString {
			delete(txt.Strings, key)
		}
	}
	for key := range txt.Integers {
		f, ok := txt.GetField(key)
		if ok == false || f.Type != TypeInt {
			delete(txt.Integers, key)
		}
	}
	for key := range txt.Bools {
		f, ok := txt.GetField(key)
		if ok == false || f.Type != TypeBool {
			delete(txt.Bools, key)
		}
	}
	for key := range txt.SliceStrings {
		f, ok := txt.GetField(key)
		if ok == false || f.Type != TypeSliceString {
			delete(txt.SliceStrings, key)
		}
	}
}
