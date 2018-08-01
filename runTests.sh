#!/bin/bash

go get -u github.com/jstemmer/go-junit-report
go test ./... -v 2>&1 | go-junit-report > report.xml

curl -X POST -d @report.xml http://localhost:9101/create
