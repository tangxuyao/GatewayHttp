package main

import (
	"github.com/gorilla/mux"
	"./v1"
	"log"
	"net/http"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("rpg-gateway"),
	)
	service.Init()

	router := mux.NewRouter().StrictSlash(true)
	v1.SetupHandler(router, service.Client())
	log.Println("listening at 8811")
	err := http.ListenAndServe(":8811", router)
	log.Fatal(err)
}
