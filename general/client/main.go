package main

import (
	"bufio"
	"log"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://localhost:8080/")

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
}
