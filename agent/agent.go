package main

import (
	"fmt"
	"log/slog"
	"os"
)

// FIXME: Clean Architecture
// TODO: Use MySQL, PostgreSQL, and SQLite sample databases

const (
	SCHEMA_FILE_PATH = "schema.sql"
)

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("failed to read file", "error", err.Error())
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return string(content), nil
}

func getData(sql string) (string, error) {
	data := sql
	return data, nil
}

func createVisualizableData(data string) (string, error) {
	visualizableData := data
	return visualizableData, nil
}

func getVisualizableData(dbType string, prompt string) (string, error) {
	var vd string

	schema, err := readFile(SCHEMA_FILE_PATH)
	if err != nil {
		return vd, err
	}

	sql, err := createQuery(dbType, schema, prompt)
	if err != nil {
		return vd, err
	}

	d, err := getData(sql)
	if err != nil {
		return vd, err
	}

	vd, err = createVisualizableData(d)
	if err != nil {
		return vd, err
	}

	return vd, nil
}
