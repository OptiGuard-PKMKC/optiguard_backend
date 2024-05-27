package main

import "net/http"

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}

	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8080", nil)
}
