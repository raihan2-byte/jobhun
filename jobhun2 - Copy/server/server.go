package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"jobhun2/dto"
	"jobhun2/logic/registration"
	"jobhun2/storage"
	"log"
	"net/http"
)

func Serve() {
	r := mux.NewRouter()
	// r := mux.NewRouter()

	mysqlStorage := &storage.MysqlStorage{}

	rm := registration.RegistrationManager{}

	//r.HandleFunc("/products", getAllProducts).Methods("GET")
	r.HandleFunc("/mahasiswa", func(w http.ResponseWriter, r *http.Request) {

		rm.Init(mysqlStorage)
		defer rm.Terminate(mysqlStorage)
		req := &dto.CreateMahasiswaRequest{}

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := rm.Register(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		re, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write(re)
	}).Methods("POST")
	//r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	r.HandleFunc("/mahasiswa/{id}", func(w http.ResponseWriter, r *http.Request) {

		rm.Init(mysqlStorage)
		defer rm.Terminate(mysqlStorage)
		req := &dto.CreateMahasiswaRequest{}

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := rm.Update(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		re, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write(re)
	}).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
