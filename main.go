package main

import (
	"fmt"
	"net/http"
	"os"
)

var versionEnvVar string

func ready(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "ready\n")
}

func healty(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "healty\n")
}

func index(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "OK\n")
}

func version(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, fmt.Sprintf("%v\n", versionEnvVar))
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	versionEnvVar = os.Getenv("VERSION")

	http.HandleFunc("/", index)
	http.HandleFunc("/healty", healty)
	http.HandleFunc("/ready", ready)
	http.HandleFunc("/version", version)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}
