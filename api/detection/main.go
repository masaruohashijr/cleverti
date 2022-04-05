package main

import (
	"cleverti/pkg/domain"
	"cleverti/pkg/service"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/anomalies", AnomalyDetection).Methods("POST")
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Detector is up on port 8080.")
	log.Fatal(srv.ListenAndServe())
}

func AnomalyDetection(w http.ResponseWriter, r *http.Request) {
	dss := domain.Datasets{}
	err := json.NewDecoder(r.Body).Decode(&dss)
	if err != nil {
		log.Println(err)
	}
	for _, ds := range dss.Datasets {
		outliers, err := service.Detect(ds.Results)
		if err == nil {
			ds.Outliers = outliers
		}
		service.CallPrinter(ds)
	}
	println(w)
}
