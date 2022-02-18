## DB

You can run a local DB using docker: `docker run --name photoaccess -e POSTGRES_PASSWORD=mysecretpassword -p 5431:5432 -d postgres`

Based on the Docker PostgreSQL documentation, the default user and db is `postgres`.  
I did not care about that in this example / demo, as this is just for a demo.

### Creation of tables

You can create the tables using an SQL script inside the `scripts` folder.

For example, in the root folder:
```bash
docker exec -i <CONTAINER_ID> psql -U <PSQL_USER> -d <DB_NAME> < scripts/create_tables.sql
```

You can also drop the tables using the `drop_tables.sql` script, in the same folder.
