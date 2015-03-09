package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/yo/", yoHandler)
	addr := fmt.Sprintf(":%v", os.Getenv("PORT"))
	log.Printf("Start server at %v", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
