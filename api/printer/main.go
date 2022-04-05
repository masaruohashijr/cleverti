package main

import (
	"cleverti/pkg/domain"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/printer", Printer).Methods("POST")
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Printer is up on port 8888.")
	log.Fatal(srv.ListenAndServe())
}

func Printer(w http.ResponseWriter, r *http.Request) {
	ds := domain.Dataset{}
	err := json.NewDecoder(r.Body).Decode(&ds)
	if err != nil {
		log.Println(err)
	}
	log.Println(ds.SiteId, ds.Outliers)
	println(w)
}
