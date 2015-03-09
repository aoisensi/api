package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func yoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("yo")
	v := url.Values{
		"username":  {os.Getenv("YO_ID")},
		"api_token": {os.Getenv("YO_API")},
	}
	resp, _ := http.PostForm("https://api.justyo.co/yo/", v)

	if resp == nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	http.Error(w, "", resp.StatusCode)
}
