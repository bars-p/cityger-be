# Migrations

## Create migration files

```
$ migrate create -ext sql -digits 3 -dir . -seq init_db
```

## Apply migrations

### General command

```
$ migrate -source file://path/to/migrations -database postgres://localhost:5432/database up 2
```

### Development setup

```
migrate -path . -database "postgres://postgres:mypassword@localhost:5000/citygerdb?sslmode=disable" -verbose up
migrate -path . -database "postgres://postgres:mypassword@localhost:5000/citygerdb?sslmode=disable" -verbose down
```

## Set Database Timezone

To update the system's default database timezone in persistent way:

```
ALTER DATABASE citygerdb SET timezone TO 'Europe/Moscow';
SELECT pg_reload_conf();
```

If `pg_reload_conf()` has no effect restart the database engine.
