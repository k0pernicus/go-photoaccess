# PhotoAccess

PhotoAccess is a test example of a Go server & RESTful API that can:
* handles photos (a bunch of `Photo` entity),
* handles annotations (linked to a `Photo` entity).

This API does not implement full CRUD for the following entity:
* `Photo`: content as base64, an ID (simple `integer`), and a creation / update date fields,
* `Annotation`: content as a string, an ID (simple `integer` too), coordinates (in pixels, so `integer` too), and a creation / update date fields.

The `Annotation` entity is linked to a `Photo` using a foreign key in DB (CASCADE to drop it if the `Photo` entity has been deleted).

## Run

There is a Makefile to build and run the service.

Basically, you could launch the service using two different commands: `make build`, which builds the Go app, and `make run`, which runs the server app.  
Please check the `Makefile` if you want to check a bit more what it does.

You can enable the debug mode using the `-debug` flag.

### Note

To get the public (only) code documentation you must install and use `godoc`: `go get golang.org/x/tools/cmd/godoc` or `go install golang.org/x/tools/cmd/godoc@latest` if you have a recent go version.  
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
docker exec -i <CONTAINER_ID> psql -U <PSQL_USER> -d <DB_NAME> < scripts/<script_name>.sql
```

You can also drop the tables using the `drop_tables.sql` script, in the same folder.

Another SQL script can be used to insert 2-3 different elements in your tables but this should not
be used anywhere else than a test environment.

## Routes

A Postman collection is available in `docs` in order to let you play easily with the API.

### Photo

For the photos, you can specify (for the `GET`) to retrieve the associated annotations or not.  
***Note***: This is managed with the query parameter `include_annotations`, that must be `true` to return all the annotations that depend of the photo(s) you are requesting.

* Get a specific photo: `/api/v1/photo/<photo_id>`, GET
* Get all photos: `/api/v1/photos`, GET
* Delete a specific photo: `/api/v1/photo/<photo_id>`, DELETE
* Create a photo: `/api/v1/photo`, POST, with the following payload:

```json
{
    "data": "..." // base64 of the image
}
```

### Annotation

* Get a specific annotation: `/api/v1/photo/<photo_id>/annotation/<annotation_id>`, GET
* Get all annotations for a photo: `/api/v1/photo/<photo_id>/annotations`, GET
* Delete a specific annotation: `/api/v1/photo/<photo_id>/annotation/<annotation_id>`, DELETE
* Create an annotation: `/api/v1/photo/<photo_id>/annotation`, POST, with the following payload:

```json
{
    "text": "this is an annotation",
    "coordinates": {
        "x": 10, // each represents the pixel (an integer)
        "x2": 20,
        "y": 10,
        "y2": 20
    }
}
```