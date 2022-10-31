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

## Regeneration from specs

The scaffolding for this serrvice was generated using the [openapi-generator](https://openapi-generator.tech) project.
You can regenerate it using that same generator e.g. via its Docker image (from repository root):

    docker run --user $(id -u):$(id -g) --rm -v$(pwd):/data -w /data openapitools/openapi-generator-cli generate -c /data/openapi-generator.yml

The generated files will all be written to the folder `internal/rest` and you will want to move the model files to `internal/rest/model`
and the following files to `internal/common`:

    auth.go
    error.go
    helpers.go
    impl.go
    logger.go
    routers.go

Make sure that the imports are updated accordingly (an IDE will usually handle that for you).

If you have updated the endpoints, you may have changes to the `internal/rest/*_service.go` files. In that case, you will
have to merge the new skeleton with your existing files because these will contain custom code.

The templates used by the generator are stored in `api/templates` and the configuration in `openapi-generator.yml`. Ignored
files are listed in `.openapi-generator-ignore`.

All custom code should either be added to the `*_service.go` files or in files outside the `rest` package and just called from the `*_service.go` files.

## Continuous Integration

The GitHub Actions CI runs [golanglint-ci](https://github.com/golangci/golangci-lint) and any unit tests it finds. PRs
are gated on both.