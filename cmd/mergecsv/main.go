package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	fileNames := getFileNames()
	var header []string
	var rows [][]string

	for _, fileName := range fileNames {
		rowData := getRowData(fileName)
		header = rowData[0]
		rows = append(rows, rowData[1:]...)
	}

	makeResultFile(header, rows)
}

func getFileNames() []string {
	files, err := os.ReadDir("./files/")
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}

func getRowData(fileName string) [][]string {
	file, err := os.Open("./files/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func makeResultFile(header []string, rows [][]string) {
	resultData := append([][]string{header}, rows...)

	resultFile, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer resultFile.Close()

	writer := csv.NewWriter(resultFile)
	defer writer.Flush()

	writer.WriteAll(resultData)
}
