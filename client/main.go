package main

import (
	"bufio"
	"errors"
	"log"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://localhost:8080/")

	if len(resp.TransferEncoding) == 0 {
		panic(errors.New("不是 chunked 编码"))
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
}
