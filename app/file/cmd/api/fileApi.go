package main

import (
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err)
	}
}
