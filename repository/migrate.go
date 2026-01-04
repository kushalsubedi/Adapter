package repository

import (
	"fmt"
	"reflect"
	"strings"
)

func goTypeToPostgres(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Int, reflect.Int64:
		return "BIGSERIAL"
	case reflect.String:
		return "TEXT"
	default:
		panic("unsupported type: " + t.String())
	}
}

func (p *PostgresRepo) AutoMigrate(model any) error {
	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("model must be a struct")
	}

	table := strings.ToLower(t.Name()) + "s"
	var columns []string

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("db")
		if tag == "" {
			continue
		}

		parts := strings.Split(tag, ",")
		col := parts[0]

		sqlType := goTypeToPostgres(f.Type)
		def := col + " " + sqlType

		if len(parts) > 1 && parts[1] == "primary" {
			def += " PRIMARY KEY"
		}

		columns = append(columns, def)
	}

	query := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (%s);",
		table,
		strings.Join(columns, ", "),
	)

	_, err := p.db.Exec(query)
	return err
}
