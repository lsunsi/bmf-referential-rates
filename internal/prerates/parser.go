package prerates

import (
  "github.com/PuerkitoBio/goquery"
  "io"
  "strconv"
  "strings"
)

// ParsedPrerate stores the delta days from today and two annualized
// rates, one for work days and one for calendar days.
type ParsedPrerate struct {
  delta   int
  rate252 float64
  rate360 float64
}

// Parse takes a reader containing an HTML page
// and parses each row into the prerate struct.
func Parse(r io.Reader) ([]ParsedPrerate, error) {
  doc, err := goquery.NewDocumentFromReader(r)

  if err != nil {
    return nil, err
  }

  prerates := []ParsedPrerate{}
  trs := doc.Find("#tb_principal1 tr:nth-child(n+3)")

  for i := range trs.Nodes {
    tds := trs.Eq(i).Find("td")

    delta, err := strconv.Atoi(tds.Eq(0).Text())
    if err != nil {
      return nil, err
    }

    rate252str := strings.TrimSpace(strings.Replace(tds.Eq(1).Text(), ",", ".", -1))
    rate252, err := strconv.ParseFloat(rate252str, 64)
    if err != nil {
      return nil, err
    }

    rate360str := strings.TrimSpace(strings.Replace(tds.Eq(2).Text(), ",", ".", -1))
    rate360, err := strconv.ParseFloat(rate360str, 64)
    if err != nil {
      return nil, err
    }

    prerates = append(prerates, ParsedPrerate{delta, rate252, rate360})
  }

  return prerates, nil
}
