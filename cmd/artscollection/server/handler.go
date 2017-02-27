package server

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/OpenGLAMTools/artscollection/collection"
	"github.com/OpenGLAMTools/artscollection/storage"
	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

var Artscollection map[string]*collection.Collection

var Storager = storage.NewTxtStorage()

// ImageMaxWidth specifies the width in pixel wich the served image is
// resized
var ImageMaxWidth = 960

// ImageMaxHeight specifies the height in pixel wich the served image is
// resized
var ImageMaxHeight = 960

// ImageResizeFunction defines the resize algorithm
var ImageResizeFunction = resize.Lanczos3

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/page/index.htm", http.StatusPermanentRedirect)
}

func allCollectionsHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Artscollection)
	errorLog(err, "allCollectionsHandler: Marshal:")
	writeBytes(b, w)
}
func collectionHandler(w http.ResponseWriter, r *http.Request) {
	coll := getCollection(r)
	err := coll.Reload()
	errorLog(err, "collectionHandler: Error reloading collection")
	b, err := coll.Marshal()
	errorLog(err, "collectionHandler: Error Marshaling coll")
	writeBytes(b, w)
}

func taxonomyHandler(w http.ResponseWriter, r *http.Request) {
	coll := getCollection(r)
	err := coll.Reload()
	errorLog(err, "collectionHandler: Error reloading collection")
	b, err := coll.Marshal()
	errorLog(err, "collectionHandler: Error Marshaling coll")
	writeBytes(b, w)
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	coll := getCollection(r)
	coll.Reload()
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
	if err != nil {
		errorLog(err, "postItemHandler: Error unmarshaling body")
		return
	}
	err = coll.WriteItem(itemID, item)
	errorLog(err, "postItemHandler: Error WriteItem")
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	coll := getCollection(r)
	coll.Reload()
	vars := mux.Vars(r)
	itemID := vars["item"]
	imgName := vars["img"]
	basepath := coll.GetBasePath()
	imgPath := filepath.Join(
		basepath,
		itemID,
		imgName,
	)
	// Resize the image
	imFile, err := os.Open(imgPath)
	errorLog(err, "imgHander: Open file:")
	img, _, err := image.Decode(imFile)
	errorLog(err, "imgHandler: Decode image:")
	t := resize.Thumbnail(uint(ImageMaxWidth), uint(ImageMaxHeight), img, ImageResizeFunction)
	// All pictures are served as png format.
	//png.Encode(w, t)
	jpeg.Encode(w, t, nil)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p := vars["page"]
	fp := filepath.Join("server", "pages", p)
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errorLog(err, "pageHandler: Error ReadFile():")
		// Send 404 and abort, when file could not read
		return
	}
	writeBytes(b, w)
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
		fmt.Fprintf(os.Stderr, "%s:%v\n", s, err)
	}
}
