package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("./Stepik/json/task_02/data.json")
	if err != nil {
		fmt.Println("Ошибка при чтении данных", err)
		return
	}

	var globalId []struct {
		Id int `json:"global_id"`
	}
	if err := json.Unmarshal(data, &globalId); err != nil {
		fmt.Println("Ошибка при парсинге:", err)
		return
	}

	counter := 0
	for _, v := range globalId {
		counter += v.Id
	}

	fmt.Printf("Result: %d\n", counter)

}
