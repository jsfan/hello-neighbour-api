FROM golang:1.18 AS build
WORKDIR /go/src
COPY internal/rest ./internal/rest
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o rest .

FROM scratch AS runtime
COPY --from=build /go/src/rest ./
EXPOSE 8080/tcp
ENTRYPOINT ["./rest"]
