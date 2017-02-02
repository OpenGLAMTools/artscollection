package server

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/OpenGLAMTools/artscollection/collection"
	"github.com/OpenGLAMTools/artscollection/storage"
	"github.com/gorilla/mux"
)

var Artscollection map[string]*collection.Collection

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
	item, ok := coll.GetItem(itemID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, err := item.Marshal()
	errorLog(err)
	writeBytes(b, w)
}

func postItemHandler(w http.ResponseWriter, r *http.Request) {
	coll := getCollection(r)
	vars := mux.Vars(r)
	itemID := vars["item"]
	rbody, err := ioutil.ReadAll(r.Body)
	errorLog(err)
	var item storage.Storager
	err = item.Unmarshal(rbody)
	errorLog(err)
	err = coll.WriteItem(itemID, item)
	errorLog(err)
}

func getCollection(r *http.Request) *collection.Collection {
	vars := mux.Vars(r)
	collID := vars["collection"]
	coll := Artscollection[collID]
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
