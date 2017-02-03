package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/OpenGLAMTools/artscollection/collection"
	"github.com/OpenGLAMTools/artscollection/storage"
	"github.com/gorilla/mux"
)

var Artscollection map[string]*collection.Collection

var Storager = storage.NewTxtStorage()

func collectionHandler(w http.ResponseWriter, r *http.Request) {
	coll := getCollection(r)
	b, err := coll.Marshal()
	errorLog(err, "collectionHandler: Error Marshaling coll")
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
	errorLog(err, "itemHandler: error at item.Marshal")
	writeBytes(b, w)
}

func postItemHandler(w http.ResponseWriter, r *http.Request) {
	coll := getCollection(r)
	vars := mux.Vars(r)
	itemID := vars["item"]
	rbody, err := ioutil.ReadAll(r.Body)
	errorLog(err, "postItemHandler: Error reading body")
	var item = Storager
	err = item.Unmarshal(rbody)
	fmt.Printf("<<%#v>>\n", item)
	if err != nil {
		errorLog(err, "postItemHandler: Error unmarshaling body")
		return
	}
	err = coll.WriteItem(itemID, item)
	errorLog(err, "postItemHandler: Error WriteItem")
}

func getCollection(r *http.Request) *collection.Collection {
	vars := mux.Vars(r)
	collID := vars["collection"]
	coll := Artscollection[collID]
	return coll
}

func writeBytes(b []byte, w http.ResponseWriter) {
	_, err := w.Write(b)
	errorLog(err, "Error writeBytes")
}

func errorLog(err error, s string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s:%v", s, err)
	}
}
