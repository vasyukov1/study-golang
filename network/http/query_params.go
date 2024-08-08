package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	var name string
	var age string
	fmt.Scan(&name)
	fmt.Scan(&age)

	baseURL := "http://127.0.0.1:8080/hello"
	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)
	fullURL := baseURL + "?" + params.Encode()

	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Response error: ", err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Reading error: ", err)
		return
	}

	fmt.Printf("%s\n", data)
}
