## Migration

**Command**: docker run --rm migrate/migrate -version (https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz)

- [x] Write RAW SQL
- [x] Create migration folder
- [x] Generate migration file _migrate create -ext sql -dir migrations -seq create_users_table_
- [x] Edit generated migration file up and down
- [x] Run: _migrate -path migrations -database "postgres://postgres:password@localhost:5432/mydb?sslmode=disable" up_

### In the case of edit a migrated file to alter a table

- [x] Generate migration file _migrate create -ext sql -dir migrations -seq create_users_table_
- [x] Edit newly generated migration file up and down
- [x] Run: _migrate -path migrations -database "postgres://postgres:password@localhost:5432/mydb?sslmode=disable" up_
