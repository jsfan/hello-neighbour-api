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
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/middlewares"
	"github.com/jsfan/hello-neighbour-api/internal/storage"
	"io/ioutil"
	"net/http"

	sw "github.com/jsfan/hello-neighbour-api/internal"

	"github.com/google/logger"
)

func main() {
	cfgFileOpt := flag.String("config", "config.yaml", "Configuration YAML file")

	flag.Parse()

	defer logger.Init("Hello Neighbour Logger", true, false, ioutil.Discard).Close()

	cfg, err := config.ReadConfig(*cfgFileOpt)
	if err != nil {
		logger.Fatalf(`Could not read configuration file "%s": %v`, *cfgFileOpt, err)
	}

	jwtKeys, err := config.ReadKeyPair(&cfg.JwtSignKeys)
	if err != nil {
		logger.Fatalf(`Could not load key pair: %v`, err)
	}

	masterStore, err := storage.ConnectStore(&cfg.Database)
	if err != nil {
		logger.Fatalf("Could not connect to database: %v", err)
	}

	if err := masterStore.Migrate(&cfg.Database.DbName); err != nil && err.Error() != "no change" {
		logger.Fatalf("Database migration failed: %v", err)
	}
	logger.Info("Server started")

	router := sw.NewRouter(jwtKeys)
	router.Use(middlewares.GetStorageMiddleWare(masterStore))
	router.Use(middlewares.AuthnMiddleware)

	logger.Fatal(http.ListenAndServe(":8080", router))
}
