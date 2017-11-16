package main

import (
	"net/http"
	"io/ioutil"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		defer r.Body.Close()

		w.Write(body)
		return
	})
	http.ListenAndServe(":8080", mux)
}
