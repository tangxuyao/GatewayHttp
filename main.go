package main

import (
	"google.golang.org/grpc"
	"github.com/gorilla/mux"
	"time"
	"golang.org/x/net/context"
	"./v1"
	"log"
	"net/http"
)

func main() {

	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//service := micro.NewService(
	//	micro.Name("rpg-gateway"),
	//)
	//service.Init()
	//
	router := mux.NewRouter().StrictSlash(true)
	v1.SetupHandler(router, ctx, conn)
	log.Println("listening at 8811")
	err = http.ListenAndServe(":8811", router)
	log.Fatal(err)
}
