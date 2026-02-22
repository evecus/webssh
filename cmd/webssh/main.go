package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/yourusername/webssh/internal/handler"
)

func main() {
	port := flag.Int("port", 8888, "HTTP server port")
	flag.Parse()

	mux := http.NewServeMux()
	handler.Register(mux)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("WebSSH server started on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
