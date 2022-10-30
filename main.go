/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	rest "github.com/jsfan/hello-neighbour-api/internal/rest"
)

func main() {
	log.Printf("Server started")

	var jwtKeys interface{}
	// TODO: Set your JWT keys here

	AdministratorApiService := rest.NewAdministratorApiService()
	AdministratorApiController := rest.NewAdministratorApiController(AdministratorApiService)

	DefaultApiService := rest.NewDefaultApiService()
	DefaultApiController := rest.NewDefaultApiController(DefaultApiService)

	LeaderApiService := rest.NewLeaderApiService()
	LeaderApiController := rest.NewLeaderApiController(LeaderApiService)

	MemberApiService := rest.NewMemberApiService()
	MemberApiController := rest.NewMemberApiController(MemberApiService)

	router := rest.NewRouter(jwtKeys, AdministratorApiController, DefaultApiController, LeaderApiController, MemberApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
