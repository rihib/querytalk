package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

type RequestBody struct {
	Prompt string `json:"prompt"`
}

type ResponseBody struct {
	VisualizableData string `json:"visualizableData"`
}

func prompt2sql(prompt string) string {
	// TODO: Implement this function
	sql := prompt
	return sql
}

func getData(sql string) string {
	// TODO: Implement this function
	data := sql
	return data
}

func createVisualizableData(data string) string {
	// TODO: Implement this function
	visualizableData := data
	return visualizableData
}

func getVisualizableData(prompt string) string {
	sql := prompt2sql(prompt)
	d := getData(sql)
	vd := createVisualizableData(d)
	return vd
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		slog.Info("Invalid method")
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	maxSize := int64(1 << 20) // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, maxSize)
	var reqb RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqb); err != nil {
		slog.Error(err.Error())

		if err.Error() == "http: request body too large" {
			http.Error(w, "Request body too large", http.StatusRequestEntityTooLarge)
			return
		}
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	vd := getVisualizableData(reqb.Prompt)
	resb := ResponseBody{
		VisualizableData: vd,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resb); err != nil {
		slog.Error(err.Error())
		return
	}
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error(err.Error())
	}
}
