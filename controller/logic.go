package main

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
