# Deall Test BE

A simple CRUD and Login JWT Token based REST API.

## Features

- CRUD Users
- Generate token access
- Refresh Token
- Validate Token
- Login Logout


## Tech

This app uses a number of open source projects to work properly:

- [Golang](https://go.dev/) - Go Language, v1.18
- [gin](https://gin-gonic.com/) - huge framework golang
- [gorm](https://gorm.io/) - awesome ORM for database
- [PostgreSQL](https://www.postgresql.org/) - great RDBMS database
- [JWT](https://jwt.io/) - token auth generator
- [Kubernetes](https://kubernetes.io/) - is an open-source system for automating deployment, scaling, and management of containerized applications


## Installation
dont forget to install [golang first](https://go.dev/dl/) 
##### - `go run main.go`
using `go run main.go` need to serve database first and set the config at `.env`, and see these db configuration, adjust them with your own :
```
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=postgres
DB_NAME=postgres
DB_PASSWORD=secret
DB_TIMEOUT=10
```

##### - [docker compose](https://docs.docker.com/compose/gettingstarted/) (recommended)
dont forget to install docker. for easy run on local docker machine. will compose 2 container: Postgresql15 and App itself
```
// for up running the app
docker-compose up

// to stop the app
docker-compose down

// issue with build (try this cmd)
docker-compose up --build
docker-compose build
```
try play something at `http://localhost:8080/api/v1`

##### - kubernetes
please follow and/or read this [guide](https://levelup.gitconnected.com/deploying-dockerized-golang-api-on-kubernetes-with-postgresql-mysql-d190e27ac09f) to deploy and run app using kubernetes locally

## Default Data
you can login with this data (generated from the beginning through migration)
```
username: admin
password: admin
```
pass the payload to `/login`
