package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://localhost:8080/")

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
}
