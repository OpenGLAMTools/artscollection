package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var ServerPort = ":8081"

func Serve() {
	router := mux.NewRouter()
	router.HandleFunc("/collection/{collection}", collectionHandler).Methods("GET")
	router.HandleFunc("/collection/{collection}/taxonomy/{term}", taxonomyHandler).Methods("GET")
	router.HandleFunc("/collection/{collection}/{item}", itemHandler).Methods("GET")
	router.HandleFunc("/collection/{collection}/{item}", postItemHandler).Methods("POST")
	fmt.Println("Starting server", ServerPort)
	err := http.ListenAndServe(ServerPort, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
