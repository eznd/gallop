#!/bin/bash
docker build . --tag=gallop:latest
docker run -p 9101:9101 -e GALLOP_PORT=9101 gallop:latest