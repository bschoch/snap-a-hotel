package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"io/ioutil"
	"net/http"
)

func main() {
	if err := InitHotelCache(); err != nil {
		panic(err)
	}
	fmt.Println("Hotel Cache Initialized", len(HotelCache))
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		defer r.Body.Close()
		var request struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Bearing   float64 `json:"bearing"`
		}
		if err := json.Unmarshal(body, &request); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println("request", request)
		hotel, err := Search(request.Latitude, request.Longitude, request.Bearing, HotelCache)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		bs, err := json.Marshal(hotel)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println("response", string(bs))
		w.Header().Set("access-control-allow-origin", "*")
		w.Write(bs)
		return
	})
	handler := cors.Default().Handler(mux)
	fmt.Println(http.ListenAndServe(":8080", handler))
}
