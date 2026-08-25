package main

import (
	"bakery34/api"
	"bakery34/service"
	"bakery34/store"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := store.Open("bakery.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	svc := service.New(db)
	mux := api.New(svc)
	fmt.Println("bakery34 listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
