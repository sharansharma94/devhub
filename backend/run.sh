#!/usr/bin/env bash
# docker build -f dockerFile . -t go-gin
# docker run -i -t -p 8080:8080 go-gin

docker-compose -f docker-compose.yml up --build