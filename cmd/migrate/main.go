package main

import (
  "github.com/golang-migrate/migrate"
  _ "github.com/golang-migrate/migrate/database/postgres"
  _ "github.com/golang-migrate/migrate/source/file"
  "log"
)

func main() {
  database := "postgres://postgres:postgres@database:5432/postgres?sslmode=disable"
  migrations := "file:///root/db"

  if m, e := migrate.New(migrations, database); e != nil {
    log.Fatal(e)
  } else if e := m.Up(); e != nil {
    log.Fatal(e)
  }
}
