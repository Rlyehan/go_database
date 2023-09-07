package models

import (
  "fmt"
  "errors"
)

type Table struct {
  Name  string
  Schema map[string]string
  Rows []Row
  Index map[string]map[interface{}][]int
}


func NewTable(name string, schema map[string]string *Table) {
  return &Table{
    Name: name,
    Schema: schema,
    Rows: []Row{},
  }
}

func (t *Table) AddRow(row Row) error {
  for colName, colType := range t.Schema {
    value, exists := row.Data[colName]
    if !exists || fmt.Sprintf("%T", value) != colType {
      return fmt.Errorf("schema validation failed for column: %s", colName)
    }
  }
  t.Rows = append(t.Rows, row)
  return nil
}

func (t *Table) CreateIndex(columnName string) error {
  if _, exists := t.Schema[columnName]; !exists {
    return fmt.Errorf("column does not exist: %s", columnName)
  }

  t,Index[columnName] = make(map[interface{}][]int)
  for i, row := range t.Rows {
    value := row.Data[columnName]
    t.Index[columnName][value] = append(t.Index[columnName][value], i)
  }
  return nil
}

fun (t *Table) QueryIndex(columnName string, value interface{}) ([]Row, error) {
  return nil
}
