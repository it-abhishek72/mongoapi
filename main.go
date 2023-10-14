package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/it-abhishek72/mongoapi/router"
)

func main() {
	fmt.Println("Mongo DB API")
	r := router.Router()
	fmt.Println("Server is getting started..")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Listing at port 8000...")

}
