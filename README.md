# Gallop

## What's this?

Deadly simple Allure-as-a-service dockerized tool

Run image, run `go test` (see below), POST to `/create` - and you get allure report link

## Requirements

 * `go get -u github.com/jstemmer/go-junit-report` - to convert go test output to junit format

## Running image

```
docker build . --tag=gallop:latest
docker run -p 9101:9101 -e GALLOP_PORT=9101 gallop:latest
```

## Creating report ang getting link

```
go get -u github.com/jstemmer/go-junit-report
go test ./... -v 2>&1 | go-junit-report > report.xml

curl -X POST -d @report.xml http://localhost:9101/create
```

If everything goes smooth you'll get HTTP 201 code and report link in `Location:` header