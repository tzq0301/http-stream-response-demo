package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		flusher, _ := w.(http.Flusher)

		s := []rune("中国人不骗中国人")

		for _, ch := range s {
			_, _ = fmt.Fprint(w, string(ch)+"\n")
			flusher.Flush() // 发一个字
			time.Sleep(500 * time.Millisecond)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
