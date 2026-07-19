package main

import (
	"log"
	"net/http"
)

func HelloFirst(w http.ResponseWriter, req *http.Request) {
	return 
}

func Web() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
