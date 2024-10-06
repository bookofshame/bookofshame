### Book Of Shame
A website to store criminal records.

### Prerequisite
- [Go](https://go.dev/dl/) (1.23.0 or later)
- [Bun](https://bun.sh) (for some script aliases)

### Environment setup
Create a `.env` file and fill in the details. See `.env-example` for reference.

Install goose for database migration.

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Database setup
> Turso CLI isn't available for windows. You can use WSL or create an account in Turso and create a DB there and use it.

Install [turso cli](https://docs.turso.tech/cli/installation). Then run the following command to start a database server.
```bash
bun db
```
After running the command above, you will see a database server running and it will give you a link to access the database like below.
```bash
http://127.0.0.1:8080
```
Add it to your `.env` file as `TURSO_DB_URL`. You don't need to add `TURSO_DB_AUTH_TOKEN`.


Now run the following command to migrate the database. See more migration commands below.
```bash
bun mg:up
```

### Run app
Once you have set up your environment following the instructions above, run the command below to start the api server.
```bash
bun start
```
The server can be accessed in this url: `http://localhost:3000`

### Hot reload (optional)
Install air for hot-reloading.
```bash
go install github.com/air-verse/air@latest
```
Hot reload server:
```bash
bun dev
```

### Migration
All the migration files are stored in `migrations` directory.

Create a new migration file.

```bash
goose -dir migrations create migration_name sql
```

Apply all available migrations

```bash
bun mg:up
```

Revert single migration from current version

```bash
bun mg:down
```

To check migration status.

```bash
bun mg:status
```

To revert all migrations

```bash
bun mg:reset
```