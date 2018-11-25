package prerates

import (
  "strings"
  "testing"
)

const html = `
  <table border="0" cellspacing="0" cellpadding="0" id="tb_principal1" name="tb_principal1" class="tabConteudo">
    <tbody>
      <tr>
        <td class="tabelaTitulo" rowspan="2" align="center">Dias<br>Corridos</td>
        <td class="tabelaTitulo" align="center" colspan="2">DI x pré</td>
      </tr>
      <tr>
        <td class="tabelaItem" align="center">252<sup>(2)(4)</sup></td>
        <td class="tabelaItem" align="center">360<sup>(1)</sup></td>
      </tr>
      <tr>
        <td class="tabelaConteudo1" align="center">3</td>
        <td nowrap="" class="tabelaConteudo1" align="center">6,40 </td>
        <td nowrap="" class="tabelaConteudo1" align="center">0,00 </td>
      </tr>
      <tr>
        <td class="tabelaConteudo2" align="center">6</td>
        <td nowrap="" class="tabelaConteudo2" align="center">6,50 </td>
        <td nowrap="" class="tabelaConteudo2" align="center">6,09 </td>
      </tr>
    </tbody>
  </table>
`

// TestParser tests the `parser` method.
func TestParser(t *testing.T) {
  prerates, err := Parse(strings.NewReader(html))

  if err != nil {
    t.FailNow()
  }

  if len(prerates) != 2 {
    t.FailNow()
  }

  if prerates[0] != (Prerate{delta: 3, rate252: 6.4, rate360: 0.0}) {
    t.FailNow()
  }

  if prerates[1] != (Prerate{delta: 6, rate252: 6.5, rate360: 6.09}) {
    t.FailNow()
  }
}
