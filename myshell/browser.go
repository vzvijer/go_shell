package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	body, err := ioutil.ReadAll(r.Body) //r.URL.Path[1:]
	if err != nil {
		fmt.Println(err)
	}
	command := string(body)
	fmt.Println(command)
	execCommand(command, nil, w)
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func startBrowserFrontend() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
