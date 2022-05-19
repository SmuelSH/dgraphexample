package main

import (
	"dgraphexample/grpcharge"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := ":3000"

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))

	})

	r.Post("/ListData", ListData())

	r.Post("/ChargeData", ChargeData())
	r.Get("/ListBuyers", ListBuyer())
	r.Post("/ListHistory", ListHistory())
	r.Post("/ListSameIP", ListSameIP())
	r.Post("/ListSugerencia", ListSugerencia())

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}

func ListData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		fmt.Println("title:" + request["title"])
		fmt.Println("past:" + request["post"])
		fmt.Println("fecha:" + request["fecha"])
		w.Write([]byte("Good job....!"))
	}
}

func ListBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		fmt.Println("title:" + request["title"])
		fmt.Println("past:" + request["post"])
		fmt.Println("fecha:" + request["fecha"])
		//fmt.Printf(grpcharge.QueryListBuyer())
		w.Write([]byte(grpcharge.QueryListBuyer()))
	}
}

func ListHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		var sidBuyer string
		json.NewDecoder(r.Body).Decode(&request)

		sidBuyer = request["idBuyer"]
		fmt.Println("idBuyer:" + sidBuyer)

		//fmt.Printf(grpcharge.QueryListHistory(sidBuyer))
		w.Write([]byte(grpcharge.QueryListHistory(sidBuyer)))
	}
}

func ListSameIP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		var sidBuyer string
		json.NewDecoder(r.Body).Decode(&request)

		sidBuyer = request["idBuyer"]
		fmt.Println("idBuyer:" + sidBuyer)

		//fmt.Printf(grpcharge.QueryListHistory(sidBuyer))
		w.Write([]byte(grpcharge.QueryIPTransac(sidBuyer)))
	}
}

func ListSugerencia() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		var sidBuyer string
		json.NewDecoder(r.Body).Decode(&request)

		sidBuyer = request["idBuyer"]
		fmt.Println("idBuyer:" + sidBuyer)

		//fmt.Printf(grpcharge.QueryListHistory(sidBuyer))
		w.Write([]byte(grpcharge.QuerySugerencia(sidBuyer)))
	}
}

func ChargeData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request := map[string]string{}
		//json.NewDecoder(r.Body).Decode(&request)
		//grpcharge.MutatedGraph()

		grpcharge.MutatedGraph()

		w.Write([]byte("Charge Ready......completed!!!!"))
	}
}

//env GO111MODULE=off go run dgraphexample.go
//env GO111MODULE=off go build dgraphexample.go
//go mod init
//go mod tidy
//require "importendpoint" v1.2.3
//replace "importendpoint" v1.2.3 => "{local path to the importendpoint module}"
