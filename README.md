# PhotoAccess

PhotoAccess is a test example of a Go server & RESTful API that can:
* handles photos (a bunch of `Photo` entity),
* handles annotations (linked to a `Photo` entity),
* handles additional photos (linked to an `Annotation` entity).

This API does not implement full CRUD for the following entity:
* `Photo`,
* `Annotation`.

## Run

There is a Makefile to build and run the service.

Basically, you could launch the service using two different commands: `make build`, which builds the Go app, and `make run`, which runs the server app.  
Please check the `Makefile` if you want to check a bit more what it does.

### Note

To get the only public documentation you must install and use `godoc`: `go get golang.org/x/tools/cmd/godoc` or `go install golang.org/x/tools/cmd/godoc@latest` if you have a recent go version.  
You can build and run the public package (`pkg`) documentation locally using `make doc`, and open your web browser at [here](http://localhost:6060/pkg/github.com/k0pernicus/go-photoaccess/).

## Configuration

A `config.yaml` file is available at the root of the project.  
You can set up the configuration of the app but also the configuration of the database.

The default configuration handles a configuration that runs 100% locally (only for tests purposes).

## About the DB

I used a PostgreSQL DB, with a Go compatibility library called [`pgx`](https://github.com/jackc/pgx), and his extension for multiple connections called `pgxpool`.

You can run a local PostgreSQL DB using docker: `docker run --name photoaccess -e POSTGRES_PASSWORD=mysecretpassword -p 5431:5432 -d postgres`

Based on the Docker PostgreSQL documentation, the default user and db is `postgres`.  
I did not care about that in this example / demo, as this is just for a demo.

### DB scripts

You can create the tables using an SQL script inside the `scripts` folder.

For example, in the root folder:
```bash
docker exec -i <CONTAINER_ID> psql -U <PSQL_USER> -d <DB_NAME> < scripts/create_tables.sql
```

You can also drop the tables using the `drop_tables.sql` script, in the same folder.
