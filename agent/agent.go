package main

import (
	"log"
	"os"
)

// FIXME: ログはslogを使うようにする。リファクタリングする。LLMを呼ぶ部分は非同期にする？

// TODO: MySQLのサンプルデータベースを使って動作確認できるplaygroundを作る。フロントエンドも作る

const (
	SCHEMA_FILE_PATH = "schema.yaml"
)

func readFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	return string(content)
}

func getData(sql string) string {
	data := sql
	return data
}

func createVisualizableData(data string) string {
	visualizableData := data
	return visualizableData
}

func getVisualizableData(prompt string) string {
	schema := readFile(SCHEMA_FILE_PATH)
	sql := createQuery(schema, prompt)
	d := getData(sql)
	vd := createVisualizableData(d)
	return vd
}
