# Movirate

Movirate is an application to apply the knowldge of gRPC in go lang by providing information about movies.

## Pre-Requizite
go 1.20
docker 
docker-compose

## How to run
first run docker-compose for database:
```bash
docker-compose up
```


then run the server by runnung the following command in the server folder:

```bash
go run ./
```
finally to consume the apis run the client from the client folder 

```bash
go run ./
```

note: the client has mini cmd UI and it is going to be updated.
