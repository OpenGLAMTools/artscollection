package server

import (
	"log"
	"net/http"

	"github.com/OpenGLAMTools/artscollection/collection"
	"github.com/gorilla/mux"
)

var artscollection map[string]collection.Collection

func collectionHandler(w http.ResponseWriter, r *http.Request) {
	coll := getCollection(r)
	b, err := coll.Marshal()
	errorLog(err)
	writeBytes(b, w)
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	coll := getCollection(r)
	vars := mux.Vars(r)
	itemID := vars["item"]
	item := coll.GetItem(itemID)
	b, err := item.Marshal()
	errorLog(err)
	writeBytes(b, w)
}

func getCollection(r *http.Request) collection.Collection {
	vars := mux.Vars(r)
	collID := vars["collection"]
	coll := artscollection[collID]
	return coll
}

func writeBytes(b []byte, w http.ResponseWriter) {
	_, err := w.Write(b)
	errorLog(err)
}
func errorLog(err error) {
	if err != nil {
		log.Println(err)
	}
}

func saveItemHandler(w http.ResponseWriter, r *http.Request) {

}
