package main

import (
	"log"
	"net/http"
	"./v1"
	"github.com/micro/go-micro"
	"github.com/gorilla/mux"
)

func main() {
	service := micro.NewService(
		micro.Name("rpg-gateway"),
	)

	router := mux.NewRouter().StrictSlash(true)
	v1.SetupHandler(router, service.Client())
	log.Println("listening at 8811")
	err := http.ListenAndServe(":8811", router)
	log.Fatal(err)
}
