package term_table

import (
  "errors"
  "strings"
  "fmt"
)

type TermTable struct {
  Header *Header
  Rows []*Row
}

type Header []string

type Row []string

func NewTermTable(header...string) *TermTable {
  var table TermTable

  h := Header(header)
  table.Header = &h

  return &table
}

func (table *TermTable) WriteRow(row...string) error {
  if len(row) != table.Columns() {
    return errors.New("Number of columns in row does not match header")
  }

  r := Row(row)
  table.Rows = append(table.Rows, &r)

  return nil
}

func (table *TermTable) Columns() int {
  return len(*table.Header)
}

func (table *TermTable) MaxColumnWidth(columnNumber int) int {
  if columnNumber >= table.Columns() {
    return 0
  }

  maxWidth := len((*table.Header)[columnNumber])

  for _, row := range table.Rows {
    if len((*row)[columnNumber]) > maxWidth {
      maxWidth = len((*row)[columnNumber])
    }
  }

  return maxWidth
}

func (table *TermTable) PadColumn(columnString string, columnNumber int) string {
  maxWidth := table.MaxColumnWidth(columnNumber)
  currentWidth := len(columnString)
  padWidth := maxWidth - currentWidth

  return columnString + strings.Repeat(" ", padWidth)
}

func (table *TermTable) PrintHorizontalBorder(columnNumber int) {
  fmt.Printf(strings.Repeat("-", table.MaxColumnWidth(columnNumber)+2))
}

func (table *TermTable) PrintElementContents(columnString string, columnNumber int) {
  fmt.Printf(" %s ", table.PadColumn(columnString, columnNumber))
}

func (table *TermTable) PrintHorizontalBorderSeperator() {
  fmt.Printf("+")
}
func (table *TermTable) PrintVerticalBorderSeperator() {
  fmt.Printf("|")
}

func (table *TermTable) PrintBorderLine() {
  table.PrintHorizontalBorderSeperator()
  for columnNumber, _ := range *table.Header {
    table.PrintHorizontalBorder(columnNumber)
    table.PrintHorizontalBorderSeperator()
  }
  fmt.Printf("\n")
}

func (table *TermTable) PrintLine(columns []string) {
  table.PrintVerticalBorderSeperator()
  for i, column := range columns {
    table.PrintElementContents(column, i)
    table.PrintVerticalBorderSeperator()
  }
  fmt.Printf("\n")
}

func (table *TermTable) PrintTable() {
  table.PrintBorderLine()
  table.PrintLine([]string(*table.Header))
  table.PrintBorderLine()

  for _, row := range table.Rows {
    table.PrintLine([]string(*row))
  }

  table.PrintBorderLine()
}
