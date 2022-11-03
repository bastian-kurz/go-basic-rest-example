# Golang basic REST api example
[![CI](https://github.com/bastian-kurz/go-basic-rest-example/actions/workflows/ci.yml/badge.svg)](https://github.com/bastian-kurz/go-basic-rest-example/actions/workflows/ci.yml)

Is a basic example how you can build a simple REST API in Go with go-chi and zap.

Based on https://github.com/golangdk

## Used Packages

- https://github.com/uber-go/zap
- https://github.com/go-chi/chi
- https://github.com/go-chi/render
- https://cs.opensource.google/go/x/sync
- https://github.com/matryer/is
- https://github.com/go-playground/validator
- https://github.com/ggwhite/go-masker

## Environment Variables
In this example we use some different ENV vars:

| Name     | HOST | PORT | APP_ENV |
|----------|------|------|------|
|  Default | localhost | 8000 | develop |


## Features

- Basic RST-API with go-chi
- Singleton logger
- Request & Response Examples
- validation against post requests
- mask of personal information in response

## Prerequisite

- golang >= 1.19.x

## How to run
```bash
git clone git@github.com:bastian-kurz/go-basic-rest-example.git && cd go-basic-rest-example
```
```bash
HOST=localhost PORT=8081 go run cmd/server/main.go
```
or
```bash
make start
```