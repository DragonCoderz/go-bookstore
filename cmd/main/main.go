package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/DragonCoderz/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter() //initializes r as new router that will be used to handle incoming http requests
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r) //registers the fact that r will handle requests at the route localhost:9010/
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}