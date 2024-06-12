package main

import (
	"flag"
	"guideventureapi/api"
	"log"
	"os"
)

func main() {
	listenAddr := flag.String("listenAddr", "0.0.0.0:"+os.Getenv("PORT"), "server address")
	createDummyData := flag.Bool("dummyData", false, "create set of dummy data for development")
	flag.Parse()

	server, err := api.NewServer(*listenAddr, *createDummyData)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Server running on port ", *listenAddr)
	log.Fatal(server.Start())
}
