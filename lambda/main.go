package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/maiguangyang/graphql-gorm/gen"
	"github.com/maiguangyang/graphql-gorm/src"
	"github.com/maiguangyang/graphql/events"
)

func main() {
	db := gen.NewDBFromEnvVars()

	eventController, err := events.NewEventController()
	if err != nil {
		panic(err)
	}

	handler := gen.GetHTTPServeMux(src.New(db, &eventController), db)
	algnhsa.ListenAndServe(handler, nil)
}
