# Hello Neighbour API
Backend for the Hello Neighbour app started at COVID-19 Global Church Hackathon.
The [prototype](https://devpost.com/software/hello-neighbour-he4ojl) won the hackathon's [Community Award](https://medium.com/faithtech/covid-19-global-church-hackathon-results-e41b23803c28).

This prototype (available [here](https://github.com/UpstreamCode/interconnect-backend)) was written in Python/Flask.
After the hackathon we decided to switch the API to Go and this repository holds the new code base.

## Development Setup
You can use an existing Postgres server as your backend or build and run a Postgres Docker container.

### Using the local Postgres container

To build and run the container, execute the following:

     docker-compose -f docker-compose.dev.yml up -d

You now have a Postgres server available on `localhost:5432`.

To start up the API, run

    go run main.go -config config.dev.yml

This will compile and run the application. To load changes, stop the API (Ctrl-C) and re-run it.

### Using an existing Postgres server

To use a local Postgres server, make a copy of `config.dev.yml` and adjust the settings in the file
to fit your setup. To compile and run the application, execute

    go run main.go -config <your config file>

where `<your config file>` is the name of your copy of the configuration file.

## Deployment

This software is not yet ready for production deployment.
