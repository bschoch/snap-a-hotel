package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	InitHotelCache()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		defer r.Body.Close()
		var hotelRequest HotelRequest
		if err := json.Unmarshal(body, &hotelRequest); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(hotelRequest)
		hotel, err := Search(hotelRequest.Latitude, hotelRequest.Longitude, hotelRequest.Bearing, HotelCache)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
		bs, err := json.Marshal(hotel)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(bs)
		return
	})
	http.ListenAndServe(":8080", mux)
}

type HotelRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Bearing   float64 `json:"bearing"`
}
