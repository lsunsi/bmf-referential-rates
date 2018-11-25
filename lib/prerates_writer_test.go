package prerates

import (
  "database/sql"
  // I'm still not sure where to put this
  // and the compiler made me add this comment.
  _ "github.com/proullon/ramsql/driver"
  "testing"
  "time"
)

// TestWriter tests the `writer` method.
func TestWriter(t *testing.T) {
  const timeTemplate = "2006-01-02"
  now, err := time.Parse(timeTemplate, "2001-01-01")
  if err != nil {
    t.FailNow()
  }

  db, err := sql.Open("ramsql", "TestWriter")
  if err != nil {
    t.FailNow()
  } else {
    defer db.Close()
  }

  migrate := "create table prerates (id serial, date date, rate252 float4, rate360 float4, imported_at timestamp)"
  if _, err := db.Exec(migrate); err != nil {
    t.FailNow()
  }

  query := "SELECT date, rate252, rate360, imported_at FROM prerates ORDER BY date ASC"
  oldRows, err := db.Query(query)
  if err != nil {
    t.FailNow()
  } else {
    defer oldRows.Close()
  }

  if oldRows.Next() {
    t.FailNow()
  }

  prerates := []ParsedPrerate{
    ParsedPrerate{delta: 3, rate252: 6.4, rate360: 0.0},
    ParsedPrerate{delta: 6, rate252: 6.5, rate360: 6.1},
  }

  Write(db, now, prerates)

  newRows, err := db.Query(query)
  if err != nil {
    t.FailNow()
  } else {
    defer newRows.Close()
  }

  matches := 0

  date1, _ := time.Parse(timeTemplate, "2001-01-04")
  date2, _ := time.Parse(timeTemplate, "2001-01-07")

  expected := []WrittenPrerate{
    WrittenPrerate{
      date: date1, rate252: 0.064, rate360: 0.0, importedAt: now},
    WrittenPrerate{
      date: date2, rate252: 0.065, rate360: 0.061, importedAt: now},
  }

  for newRows.Next() {
    var date time.Time
    var rate252 float64
    var rate360 float64
    var importedAt time.Time
    err := newRows.Scan(&date, &rate252, &rate360, &importedAt)
    if err != nil {
      t.FailNow()
    }

    target := expected[matches]
    if target.date != date ||
      target.rate252 != rate252 ||
      target.rate360 != rate360 ||
      target.importedAt != importedAt {
      t.FailNow()
    } else {
      matches++
    }
  }

  if matches != len(expected) {
    t.FailNow()
  }
}
