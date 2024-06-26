package main

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
	w.WriteHeader(http.StatusOK)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func main() {

	// GET istekelri
	http.HandleFunc("/", indexHandler)      //http://localhost:8080
	http.HandleFunc("/index", indexHandler) // http://localhost:8080/index
	http.HandleFunc("/about", aboutHandler) // http://localhost:8080/about

	http.ListenAndServe(":8080", nil)
}
