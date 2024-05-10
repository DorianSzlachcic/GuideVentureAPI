package main

import (
	"flag"
	"guideventureapi/api"
	"guideventureapi/db"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "the server address")
	flag.Parse()

	sqliteDb, err := db.NewDb()
	if err != nil {
		log.Panic(err)
	}

	err = sqliteDb.CreateDummyData()
	if err != nil {
		log.Panic(err)
	}

	server := api.NewServer(*listenAddr)
	log.Println("Server running on port ", *listenAddr)
	log.Fatal(server.Start())
}
