package main

import (
	"flag"
	"guideventure/api/api"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "the server address")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	log.Println("Server running on port ", *listenAddr)
	log.Fatal(server.Start())
}
