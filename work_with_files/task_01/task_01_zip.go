package task_01

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func Solution() {

	// Открываем zip-архив
	r, err := zip.OpenReader("./Stepik/work_with_files/task_01/task.zip")
	if err != nil {
		log.Fatalf("О`шибка при открытии архива: %v", err)
	}
	defer r.Close() // Закрываем архив после завершения работы

	// Извлекаем CSV файл из архива
	csvFilePath, csvData, err := extractCSVFromZip(r)
	if err != nil {
		log.Fatal(err)
	}

	// Выводим путь к CSV файлу и его содержимое
	fmt.Printf("Путь к CSV файлу: %s\n", csvFilePath)
	displayCSVData(csvData)
}

// extractCSVFromZip ищет и извлекает CSV файл из zip архива
func extractCSVFromZip(r *zip.ReadCloser) (string, [][]string, error) {
	for _, f := range r.File {
		// Читаем файл из архива
		data, err := readFileFromZip(f)
		if err != nil {
			return "", nil, err
		}

		// Проверяем, является ли файл CSV
		if isCSV(data) {
			// Читаем и парсим данные из CSV файла
			csvData, err := readCSVData(data)
			if err != nil {
				return "", nil, err
			}
			return f.Name, csvData, nil
		}
	}
	return "", nil, fmt.Errorf("CSV файл не найден в архиве")
}

// readFileFromZip читает файл из zip архива и возвращает его содержимое в виде байтов
func readFileFromZip(f *zip.File) ([]byte, error) {
	fileReader, err := f.Open() // Открываем файл для чтения
	if err != nil {
		return nil, fmt.Errorf("Ошибка при открытии файла: %v", err)
	}
	defer fileReader.Close() // Закрываем файл после чтения

	// Читаем содержимое файла в буфер
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, fileReader); err != nil {
		return nil, fmt.Errorf("Ошибка при чтении файла: %v", err)
	}

	return buf.Bytes(), nil // Возвращаем содержимое файла
}

// isCSV проверяет, является ли файл CSV по наличию запятых в его содержимом
func isCSV(data []byte) bool {
	return strings.Contains(string(data), ",")
}

// readCSVData читает и парсит данные из CSV файла
func readCSVData(data []byte) ([][]string, error) {
	reader := csv.NewReader(bytes.NewReader(data)) // Создаем новый CSV ридер
	return reader.ReadAll()                        // Читаем все данные из CSV файла
}

// displayCSVData выводит данные из CSV файла на экран
func displayCSVData(data [][]string) {
	fmt.Println("Данные из CSV файла:")
	if len(data) == 0 {
		fmt.Println("Файл CSV пуст.")
		return
	}

	// Выводим каждую строку CSV файла
	for i, row := range data {
		fmt.Printf("Строка %d: %v\n", i+1, row)
	}

	// Проверяем, что в данных есть нужная строка и столбец
	if len(data) > 4 && len(data[4]) > 2 {
		fmt.Printf("Результат: %s\n", data[4][2])
	} else {
		log.Fatal("Нет данных в ожидаемом месте")
	}
}
