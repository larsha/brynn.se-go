package main

import (
	"log"
	"net/http"
	"os"

	"github.com/larsha/brynn.se-go/app/route"
)

func main() {
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), route.Load()))
}
