package collection

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/OpenGLAMTools/artscollection/storage"
)

// StorageFile is the default name
var StorageFile = "data.json"

// FieldsConfFile defines the default values for the storages
// That file has to be set into the root folder of the collection
var FieldsConfFile = "conf.yaml"

// SupportedImageExt defines the extensions for the supported Images
// If a file with such extension is located inside the item folder it
// will be handled as an image.
var SupportedImageExt = []string{".jpg", ".jpeg", ".png"}

// Item is used as API to provide also the images of the collection inside
// of one object.
type Item struct {
	ID     string       `json:"-"`
	Data   *storage.Txt `json:"data"`
	Title  string       `json:"title"`
	Images []string     `json:"images"`
}

func NewItem(s *storage.Txt, id string) *Item {
	var item Item
	item.ID = id
	item.Data = s
	item.SetTitle()
	return &item
}

func (i *Item) SetTitle() {
	i.Title, _ = i.Data.GetString("title")
	if i.Title == "" {
		i.Title = i.ID
	}
}

// Marshal is used inside the httphandler to marshal the item
func (i *Item) Marshal() ([]byte, error) {
	i.SetTitle()
	return json.MarshalIndent(i, "", "  ")
}

// Collection defines the fields for a storage that fields are used as
// default for new storages.
type Collection struct {
	// Default fields for all the storages of that collection
	Fields []storage.Field
	// The data storage
	//Storages map[string]*storage.Txt
	// Items include the storages
	Items map[string]*Item
	// All the images from a folder
	Images map[string][]string
	// Index map for the SliceString values. The logik of the keys is
	// map[Field.Key]map[SliceStringValue]
	// SliceStrings are used to support taxonomy like for example tags
	SliceStringIndex map[string]map[string][]*Item
	// The filepath which is the root folder of the collection.
	fpath string
}

// NewCollection returns an empty collection
func NewCollection(fpath string) *Collection {
	return &Collection{
		Items:            make(map[string]*Item),
		SliceStringIndex: make(map[string]map[string][]*Item),
		Images:           make(map[string][]string),
		fpath:            fpath,
	}
}

// GetBasePath returns the filepath to the collection inside of the filesystem
func (c *Collection) GetBasePath() string {
	return c.fpath
}

// GetItem returns a storage from the collection
func (c *Collection) GetItem(ID string) (Item, bool) {
	i, ok := c.Items[ID]
	i.Images = c.Images[ID]
	return *i, ok
}

func (c *Collection) WriteItem(ID string, s *storage.Txt) error {
	itemPath := filepath.Join(
		c.Path(),
		ID,
		StorageFile,
	)
	return storage.WriteTxt(s, itemPath)
}

// GetItemImages returns the image names, the basepath to load the
// images from the filesystem and a ok value
func (c *Collection) GetItemImages(ID string) (images []string, basepath string, ok bool) {
	images, ok = c.Images[ID]
	return images, c.fpath, ok
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
				c.SliceStringIndex[f.ID()] = make(map[string][]*Item)
			}
			for _, s := range c.Items {
				vals, ok := s.Data.GetSliceString(f.ID())
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
		c.Items[fi.Name()] = NewItem(s, fi.Name())
		//c.Storages[fi.Name()] = s
		c.Images[fi.Name()], err = loadImages(filepath.Join(fpath, fi.Name()))
		if err != nil {
			return c, err
		}

	}
	c.MakeSliceStringIndex()
	return c, nil
}

func loadImages(fpath string) ([]string, error) {
	var images []string
	dir, err := ioutil.ReadDir(fpath)
	if err != nil {
		return images, err
	}
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		imgPath := fi.Name()
		if isSupportedImage(imgPath) {
			images = append(images, imgPath)
		}
	}
	return images, nil
}

// returns true, when the file is a supported image
func isSupportedImage(fpath string) bool {
	for _, e := range SupportedImageExt {
		if strings.EqualFold(e, filepath.Ext(fpath)) {
			return true
		}
	}
	return false
}
