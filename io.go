package gotools

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
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

// Создание нового .CSV документа и вывод данных в него
func CreateAndWriteCSV(fileName string, headers []string, values [][]string) (err error) {
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

// Создание нового .XLSX документа и вывод данных в него
func CreateAndWriteXLSX(fileName string, headers []string, values [][]interface{}) (err error) {
	f := xlsx.New()
	defer f.Close()

	s := f.AddSheet("sheet 1")

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

	err = f.SaveAs(fileName)
	if err != nil {
		return
	}

	return
}

// Запись в существующий .XLSX документ
func WriteXLSX(w *io.Writer, headers []string, values [][]interface{}) (err error) {
	var f *xlsx.Spreadsheet
	f, err = xlsx.Open(w)
	if err != nil {
		return
	}
	defer f.Close()

	s := f.Sheet(0)

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

	err = f.Save()
	if err != nil{
		return
	}

	return
}

// Запись в существующий .XLSX документ значения в определенную ячейку
func WriteXLSXValue(w *io.Writer, row int, col int, value interface{}) (err error) {
	var f *xlsx.Spreadsheet
	f, err = xlsx.Open(w)
	if err != nil {
		return
	}
	defer f.Close()

	f.Sheet(0).Cell(col, row).SetValue(value)
	return
}
