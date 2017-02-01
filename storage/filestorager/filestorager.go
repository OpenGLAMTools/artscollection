package filestorager

import "errors"
import "github.com/OpenGLAMTools/artscollection/storage"

type Txt struct {
	Fields   []storage.Field
	Strings  map[string]string
	Integers map[string]int
	Bools    map[string]bool
}

var ErrTypeNotSupported = errors.New("Type is not supported!")

func (txt *Txt) Set(fieldID string, value interface{}) error {
	switch t := value.(type) {
	case string:
		txt.Strings[fieldID] = t
	case int:
		txt.Integers[fieldID] = t
	case bool:
		txt.Bools[fieldID] = t
	default:
		return ErrTypeNotSupported
	}
	return nil
}

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

// GetField returns the field
func (txt *Txt) GetField(fieldID) (storage.Field, bool) {

}

// AddField adds a field to the storage
func (txt *Txt) AddField(f storage.Field) error {
	txt.Fields[f.ID()] = f
	return nil
}
