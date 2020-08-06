package gotools

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/plandem/xlsx"
)

// Построчное чтение из файла в массив
func ReadFileLineByLine(fileName string) (ids []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

// Вывод данных в CSV
func OutputCSV(fileName string, headers []string, values [][]string) (err error) {
	var file *os.File
	file, err = os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := csv.NewWriter(file)
	err = w.Write(headers)
	if err != nil {
		return
	}

	for i := range values {
		err = w.Write(values[i])
		if err != nil {
			return
		}
	}
	w.Flush()
	return
}

// Вывод данных в XLSX
func OutputXSLX(fileName string, headers []string, values [][]interface{}) (err error) {
	file := xlsx.New()
	defer file.Close()

	s := file.AddSheet("sheet 1")

	// Выставляем заголовки
	for i := range headers {
		err = s.Col(i).Cell(0).SetText(headers[i])
		if err != nil {
			return
		}
	}

	// Записываем данные
	for i := range values {
		for k := range values[i] {
			s.Col(k).Cell(i + 1).SetValue(values[i][k])
		}
	}

	err = file.SaveAs(fileName)
	if err != nil {
		return
	}

	return
}
