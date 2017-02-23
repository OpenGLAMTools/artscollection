package collection

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/OpenGLAMTools/artscollection/storage"
)

// StorageFile is the default name
var StorageFile = "data.json"

// FieldsConfFile defines the default values for the storages
// That file has to be set into the root folder of the collection
var FieldsConfFile = "conf.yaml"

// Collection defines the fields for a storage that fields are used as
// default for new storages.
type Collection struct {
	// Default fields for all the storages of that collection
	Fields []storage.Field
	// The data storage
	Storages map[string]*storage.Txt
	// Index map for the SliceString values. The logik of the keys is
	// map[Field.Key]map[SliceStringValue]
	// SliceStrings are used to support taxonomy like for example tags
	SliceStringIndex map[string]map[string][]*storage.Txt
	// The filepath which is the root folder of the collection.
	fpath string
}

// NewCollection returns an empty collection
func NewCollection(fpath string) *Collection {
	return &Collection{
		Storages:         make(map[string]*storage.Txt),
		SliceStringIndex: make(map[string]map[string][]*storage.Txt),
		fpath:            fpath,
	}
}

// GetItem returns a storage from the collection
func (c *Collection) GetItem(ID string) (*storage.Txt, bool) {
	i, ok := c.Storages[ID]
	return i, ok
}

func (c *Collection) WriteItem(ID string, s *storage.Txt) error {
	itemPath := filepath.Join(
		c.Path(),
		ID,
		StorageFile,
	)
	return storage.WriteTxt(s, itemPath)
}

// Marshal returns the whole collection as json object
func (c *Collection) Marshal() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

// Path returns the root path to the collection
func (c *Collection) Path() string {
	return c.fpath
}

// Reload the collection
func (c *Collection) Reload() error {
	cnew, err := LoadTxt(c.fpath)
	if err != nil {
		return err
	}

	*c = *cnew
	//fmt.Printf("%v\n\n", cnew)
	return nil
}

// MakeSliceStringIndex creates the index for all the SliceString values
func (c *Collection) MakeSliceStringIndex() {
	for _, f := range c.Fields {
		if f.Type == storage.TypeSliceString {
			// If map is not allocated for a field
			if c.SliceStringIndex[f.ID()] == nil {
				c.SliceStringIndex[f.ID()] = make(map[string][]*storage.Txt)
			}
			for _, s := range c.Storages {
				vals, ok := s.GetSliceString(f.ID())
				if ok {
					for _, v := range vals {
						c.SliceStringIndex[f.ID()][v] = append(c.SliceStringIndex[f.ID()][v], s)
					}
				}
			}
		}
	}
}

// LoadTxt the collection from a given path
func LoadTxt(fpath string) (*Collection, error) {
	c := NewCollection(fpath)
	confFile := filepath.Join(
		fpath,
		FieldsConfFile,
	)
	var err error
	c.Fields, err = storage.LoadFields(confFile)
	if err != nil {
		return c, err
	}
	dir, err := ioutil.ReadDir(fpath)
	if err != nil {
		return c, err
	}
	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		storagePath := filepath.Join(
			fpath,
			fi.Name(),
			StorageFile,
		)
		s := storage.NewTxtStorage()
		err = storage.LoadTxt(storagePath, s)
		if err != nil {
			return c, err
		}
		if len(s.Fields) == 0 {
			s.Fields = c.Fields
		}
		c.Storages[fi.Name()] = s
	}
	c.MakeSliceStringIndex()
	return c, nil
}
