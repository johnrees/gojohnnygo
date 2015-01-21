package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", gojohnnygo)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func gojohnnygo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "<h1>go johnny go!</h1>")
	fmt.Fprintln(res, "<p>A very simple website served by golang. <a href='https://github.com/johnrees/gojohnnygo'>Github Source</b></p>")
}
