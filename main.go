package main

import (
	"log"
	"net/http"
	"os"

	"github.com/maiguangyang/graphql/events"
	// "github.com/rs/cors"
	"github.com/maiguangyang/graphql-gorm/gen"
	"github.com/maiguangyang/graphql-gorm/src"
)

const (
	defaultPort = "80"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := gen.NewDBFromEnvVars()
	defer db.Close()
	db.AutoMigrate()

	eventController, err := events.NewEventController()
	if err != nil {
		panic(err)
	}

	mux := gen.GetHTTPServeMux(src.New(db, &eventController), db)

	mux.HandleFunc("/healthcheck", func(res http.ResponseWriter, req *http.Request) {
		if err := db.Ping(); err != nil {
			res.WriteHeader(400)
			res.Write([]byte("ERROR"))
			return
		}
		res.WriteHeader(200)
		res.Write([]byte("OK"))
	})

	handler := mux
	// use this line to allow cors for all origins/methods/headers (for development)
	// handler := cors.AllowAll().Handler(mux)

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
