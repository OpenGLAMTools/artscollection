package collection

import (
	"github.com/OpenGLAMTools/artscollection/storage"
)

type TxtCollection struct {
	Fields   []*storage.Field
	Storages []*storage.Txt
}
