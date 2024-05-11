package main

import (
	"flag"
	"guideventureapi/api"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "the server address")
	flag.Parse()

	server, err := api.NewServer(*listenAddr)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Server running on port ", *listenAddr)
	log.Fatal(server.Start())
}
