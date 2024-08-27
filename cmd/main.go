package main

import (
	"log"

	"github.com/giomhern/ecomm/cmd/api"
)

func main(){
	server := api.NewAPIServer(":8080", nil)

	if error := server.Run(); error != nil {
		log.Fatal("Error found", error)
	}
}