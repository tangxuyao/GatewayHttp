package main

import (
	"log"
	"net/http"
	"./v1"
	"os"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro"
	"github.com/gorilla/mux"
)

func main() {
	host := os.Getenv("consul.host")
	log.Printf("[consul] registry at %s\n", host)

	service := micro.NewService(
		micro.Registry(registry.NewRegistry(registry.Addrs(host))),
		micro.Name("account"),
	)

	router := mux.NewRouter().StrictSlash(true)
	v1.SetupHandler(router, service.Client())
	log.Println("listening at 8811")
	err := http.ListenAndServe(":8811", router)
	log.Fatal(err)
}
