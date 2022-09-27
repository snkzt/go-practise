package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// Issue an HTTP Get request to a server.
	// http.Get creates and http.Client object
	response, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)
	scanner := bufio.NewScanner(response.Body)
	// Print the first 5 lines of the response body.
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
