package gotools

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
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