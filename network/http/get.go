package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:5555/get")
	if err != nil {
		fmt.Println("Response error:", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading error: ", err)
		return
	}

	fmt.Printf("%s\n", data)
}
