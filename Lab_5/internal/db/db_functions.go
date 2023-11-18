package db

import (
	"database/sql"
	"fmt"
)

type Database interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

type Service struct {
	DB Database
}

func New(db Database) Service {
	return Service{DB: db}
}

func (service Service) GetNames() ([]string, error) {
	query := "SELECT name FROM users"

	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error while executing request: %w", err)
	}
	defer rows.Close()

	var names []string
	// comment fix this
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("Error while scanning string: %w", err)
		}
		// comment fix this
		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error when iterating over rows: %w", err)
	}

	return names, nil
}

func (service Service) SelectUniqueValues(columnName string, tableName string) ([]string, error) {
	query := fmt.Sprintf("SELECT DISTINCT %s FROM %s", columnName, tableName)

	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error while executing request: %w", err)
	}
	defer rows.Close()

	var values []string
	// comment fix this
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, fmt.Errorf("Error while scanning string: %w", err)
		}
		// comment fix this
		values = append(values, value)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error when iterating over rows: %w", err)
	}

	return values, nil
}
