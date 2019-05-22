package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port != "" {
		port = ":" + port
	} else {
		port = ":80"
	}
	version := os.Getenv("VERSION")

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		response := "env VERSION is not set.\n\ndocker run -d \\\n\t-p 3000:80 \\\n\t-e VERSION=1.0.0 \\\n\ttheremix/http-env-server"
		if version != "" {
			response = fmt.Sprintf("HTTP env server. version:%s", version)
		}
		io.WriteString(w, response)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
