/*
 * Hello Neighbour
 *
 * This is the API for the 'Hello Neighbour' project inspired from the COVID-19 Global Church Hack
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 * was /main.go
 */

package main

import (
	"flag"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/storage"
	"log"
	"net/http"

	sw "github.com/jsfan/hello-neighbour/internal"
)

func main() {

	cfgFileOpt := flag.String("config", "config.yaml", "Configuration YAML file")

	flag.Parse()

	cfg, err := config.ReadConfig(*cfgFileOpt)
	if err != nil {
		log.Fatalf(`[ERROR] Could not read configuration file "%s": %v`, *cfgFileOpt, err)
	}

	jwtKeys, err := config.ReadKeyPair(&cfg.JwtSignKeys)
	if err != nil {
		log.Fatalf(`[ERROR] Could not load key pair: %v`, err)
	}

	_, err = storage.Connect(&cfg.Database)
	if err != nil {
		log.Fatalf("[ERROR] Could not connect to database: %v", err)
	}

	log.Printf("Server started")

	router := sw.NewRouter(jwtKeys)
	router.Use(sw.AuthMiddleware)

	log.Fatal(http.ListenAndServe(":8080", router))
}
