package main

import (
	"flag"
	"guideventureapi/api"
	"guideventureapi/db/sqlite"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	listenAddr := flag.String("listenAddr", ":"+port, "server address")
	createDummyData := flag.Bool("dummyData", false, "create set of dummy data for development")
	flag.Parse()

	sqliteDb, err := sqlite.NewSQLiteDb()
	if err != nil {
		log.Panic(err)
	}

	serverOptions := []api.Option{api.WithListenAddr(*listenAddr)}
	if *createDummyData {
		serverOptions = append(serverOptions, api.WithDummyData())
	}
	server, err := api.NewServer(sqliteDb, serverOptions...)
	if err != nil {
		log.Panic(err)
	}

	log.Fatal(server.Start())
}
