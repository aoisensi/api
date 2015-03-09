package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	initStatus()
}

func main() {
	http.HandleFunc("/yo/", yoHandler)
	http.HandleFunc("/status/", statusHandler)
	addr := fmt.Sprintf(":%v", os.Getenv("PORT"))
	log.Printf("Start server at %v", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func setup(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
