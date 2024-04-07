package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// FIXME: Clean Architecture
// TODO: Use MySQL and PostgreSQL sample databases

const (
	SCHEMA_FILE_PATH = "schema.sql"
	DB_FILE_PATH     = "chinook.db"
)

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("failed to read file", "error", err.Error())
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return string(content), nil
}

func execQuery(query string) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", DB_FILE_PATH)
	if err != nil {
		slog.Error("failed to open database", "msg", err.Error())
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		slog.Error("failed to query database", "msg", err.Error())
		return nil, fmt.Errorf("failed to query database: %v", err)
	}

	return rows, nil
}

func createVisualizableData(rows *sql.Rows, data string) (string, error) {
	columns, err := rows.Columns()
	if err != nil {
		slog.Error("failed to get columns", "msg", err.Error())
		return "", fmt.Errorf("failed to get columns: %v", err)
	}

	result := make(map[string]interface{})
	result["data"] = []map[string]interface{}{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := 0; i < len(columns); i++ {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			slog.Error("failed to scan rows", "msg", err.Error())
			return "", fmt.Errorf("failed to scan rows: %v", err)
		}

		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			var value interface{}
			val := values[i]

			b, ok := val.([]byte)
			if ok {
				value = string(b)
			} else {
				value = val
			}

			rowMap[colName] = value
		}

		result["data"] = append(result["data"].([]map[string]interface{}), rowMap)
	}

	if err := rows.Err(); err != nil {
		slog.Error("failed to get rows", "msg", err.Error())
		return "", fmt.Errorf("failed to get rows: %v", err)
	}
	defer rows.Close()

	var dataObj map[string]interface{}
	err = json.Unmarshal([]byte(data), &dataObj)
	if err != nil {
		slog.Error("failed to unmarshal JSON", "msg", err.Error())
		return "", fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	result["chart"] = dataObj

	jsonResult, err := json.Marshal(result)
	if err != nil {
		slog.Error("failed to marshal JSON", "msg", err.Error())
		return "", fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return string(jsonResult), nil
}

func getVisualizableData(dbType string, prompt string) (string, error) {
	schema, err := readFile(SCHEMA_FILE_PATH)
	if err != nil {
		return "", err
	}

	query, data, err := createQueryAndData(dbType, schema, prompt)
	if err != nil {
		return "", err
	}

	rows, err := execQuery(query)
	if err != nil {
		return "", err
	}

	vd, err := createVisualizableData(rows, data)
	if err != nil {
		return "", err
	}

	return vd, nil
}
