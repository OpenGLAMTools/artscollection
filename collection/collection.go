package collection

import (
	"github.com/OpenGLAMTools/artscollection/storage"
)

// Collection is the interface for handling a lot of Items, which are representated
// over a storage.
type Collection interface {
	//GetItemList() []Item
	GetItem(ID string) Item
	//SetItem(ID string, item Item)
	//DeleteItem(ID string)
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type Item interface {
	storage.Storager
}
