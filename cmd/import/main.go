package main

import (
  "bmf-referential-rates/internal/prerates"
  "database/sql"
  _ "github.com/lib/pq"
  "log"
  "net/http"
  "time"
)

const dbURL = "postgres://postgres:postgres@database:5432/postgres?sslmode=disable"
const preratesURL = "http://www2.bmf.com.br/pages/portal/bmfbovespa/boletim1/TxRef1.asp"

func main() {
  resp, err := http.Get(preratesURL)
  if err != nil {
    log.Fatal(err)
  }

  prs, err := prerates.Parse(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  db, err := sql.Open("postgres", dbURL)
  if err != nil {
    log.Fatal(err)
  }

  if err := prerates.Write(db, time.Now(), prs); err != nil {
    log.Fatal(err)
  }
}
