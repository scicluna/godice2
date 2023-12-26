package main

import (
	"godice/rendering"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", rendering.ServeTemplate)

	http.ListenAndServe(":8080", nil)
}
