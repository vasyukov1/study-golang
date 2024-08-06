package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type group struct {
	ID       int
	Number   string
	Year     int
	Students []student
}

type student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type result struct {
	Average float64
}

func main() {

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Ошибка при чтении даннных:", err)
		return
	}

	//data := []byte(`
	//{
	//"id": 1,
	//"number": "123",
	//"year": 2021,
	//"students": [
	//    {
	//        "last_name": "Ivanov",
	//        "first_name": "Ivan",
	//        "middle_name": "Ivanovich",
	//        "birthday": "2000-01-01",
	//        "address": "Some address",
	//        "phone": "1234567890",
	//        "rating": [4, 5, 3, 5]
	//    },
	//    {
	//        "last_name": "Petrov",
	//        "first_name": "Petr",
	//        "middle_name": "Petrovich",
	//        "birthday": "2000-02-02",
	//        "address": "Another address",
	//        "phone": "0987654321",
	//        "rating": [5, 5, 5, 4]
	//    }
	//]
	//}`)

	var g group
	if err := json.Unmarshal(data, &g); err != nil {
		fmt.Println("Ошибка при парсинге:", err)
		return
	}
	countStudents := len(g.Students)
	countRatings := 0
	for i := 0; i < countStudents; i++ {
		countRatings += len(g.Students[i].Rating)
	}
	averageRating := float64(countRatings) / float64(countStudents)

	res := result{averageRating}

	answer, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при создании JSON:", err)
		return
	}
	fmt.Printf("%s\n", answer)
}
