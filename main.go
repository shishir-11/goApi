package main

import (
	"fmt"
	"go-postgres/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Start Listening on Port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
