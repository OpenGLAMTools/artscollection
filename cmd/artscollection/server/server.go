package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/as27/golib/css/semanticcss227"
	"github.com/as27/golib/js/jquerymin"
	"github.com/as27/golib/js/semanticjs227"
	"github.com/as27/golib/js/vuejsdev"
	"github.com/as27/golib/js/vueresourcemin"
	"github.com/as27/golib/js/vueroutermin"
	"github.com/gorilla/mux"
)

import _ "net/http/pprof"

var ServerPort = ":8081"

func Serve() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/collection", allCollectionsHandler).Methods("GET")
	router.HandleFunc("/collection/{collection}", collectionHandler).Methods("GET")

	router.HandleFunc("/collection/{collection}/taxonomy/{term}", taxonomyHandler).Methods("GET")
	router.HandleFunc("/collection/{collection}/{item}", itemHandler).Methods("GET")
	router.HandleFunc("/collection/{collection}/{item}", postItemHandler).Methods("POST")
	router.HandleFunc("/collection/{collection}/{item}/{img}", imgHandler).Methods("GET")
	router.HandleFunc("/page/{page}", pageHandler).Methods("GET")
	router.HandleFunc("/lib/css/semantic.min.css", semanticcss227.Handler).Methods("GET")

	router.HandleFunc("/lib/js/vue.min.js", vuejsdev.Handler).Methods("GET")
	router.HandleFunc("/lib/js/vue-resource.min.js", vueresourcemin.Handler).Methods("GET")
	router.HandleFunc("/lib/js/vue-router.min.js", vueroutermin.Handler).Methods("GET")
	router.HandleFunc("/lib/js/jquery.min.js", jquerymin.Handler).Methods("GET")
	router.HandleFunc("/lib/js/semantic.js", semanticjs227.Handler).Methods("GET")

	//router.HandleFunc("/lib/js/react.min.js", reactmin.Handler).Methods("GET")
	//router.HandleFunc("/lib/js/react-dom.min.js", reactdommin.Handler).Methods("GET")
	//router.HandleFunc("/lib/js/react-jsonschema-form.js", reactjsonschemaform.Handler).Methods("GET")

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("Starting server", ServerPort)
	err := http.ListenAndServe(ServerPort, router)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
