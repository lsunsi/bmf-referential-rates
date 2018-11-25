package prerates

import (
  "database/sql"
  "time"
)

const statement = `
  insert into prerates (date, rate252, rate360, imported_at)
  values ($1, $2, $3, $4)
`

// WrittenPrerate represents the prerate data actually in the database
type WrittenPrerate struct {
  date       time.Time
  rate252    float64
  rate360    float64
  importedAt time.Time
}

// Write takes care of actually writing the parsed
// prerates values into the database.
func Write(db *sql.DB, now time.Time, prerates []ParsedPrerate) error {
  for _, prerate := range prerates {
    date := now.AddDate(0, 0, prerate.delta)
    _, err := db.Exec(statement, date, prerate.rate252/100, prerate.rate360/100, now)
    if err != nil {
      return err
    }
  }
  return nil
}
