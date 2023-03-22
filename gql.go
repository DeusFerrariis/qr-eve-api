package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"	
)

func GenerateSchema() {
	attendeeType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Event",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"fullName": &graphql.Field{
					Type: graphql.String,
				},
				"anonymous": &graphql.Field{
					Type: graphql.Bool,
				},
			}
		}
	)

	eventType:= graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Event",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"date": &graphql.Field{
					Type: graphql.String,
				},
				"location": &graphql.Field{
					Type: graphql.String,
				},
				"anonAllowed": &graphql.Field{
					Type: graphql.Bool,
				},
				"attendees": &graphql.Field{
					Type: graphql.NewList(attendeeType)
				},
			}
		}
	)


	queryFields := graphql.Fields{
		"event": &graphql.Field{
			Type: eventType,
			Resolve: 
		}
	}
}
