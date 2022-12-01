package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const MAIN_API = "https://v2.jokeapi.dev/joke/"

func getJoke(w http.ResponseWriter) {
	response, err := http.Get(MAIN_API + "any")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(responseData))
}

func main() {

	http.HandleFunc("/joke", jokeHandler)

	fmt.Printf("Starting server at port 8090\n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/joke" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	getJoke(w)

}
