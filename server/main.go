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

		for i := 1; i <= 10; i++ {
			_, _ = fmt.Fprintf(w, "Chunk #%d\n", i)
			flusher.Flush() // Trigger "chunked" encoding and send a chunk...
			time.Sleep(500 * time.Millisecond)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
