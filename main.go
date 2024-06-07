package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	status := mainExit()
	if status != 0 {
		// From os/proc.go: "For portability, the status code should be in the range [0, 125]."
		if status < 0 || status > 125 {
			status = 125
		}
		os.Exit(status)
	}
}

func mainExit() int {
	port := ":3000"

	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Print request details
		fmt.Printf("+++ Received request\n")
		fmt.Printf("Host: %s\n", r.Host)
		fmt.Printf("Method: %s\n", r.Method)
		fmt.Printf("Proto: %s\n", r.Proto)
		fmt.Printf("RemoteAddr: %s\n", r.RemoteAddr)
		fmt.Printf("URL: %s\n", r.URL.String())
		fmt.Printf("RequestURI: %s\n", r.RequestURI)
		fmt.Printf("Endpoint: %s\n", r.URL.Path)

		fmt.Printf("Headers:\n")
		for key, values := range r.Header {
			for _, value := range values {
				fmt.Printf("  %s: %s\n", key, value)
			}
		}

		// Print query parameters
		queryParams := r.URL.Query()
		fmt.Printf("Query Params:\n")
		for key, values := range queryParams {
			for _, value := range values {
				fmt.Printf("  %s: %s\n", key, value)
			}
		}

		// Print request body
		fmt.Printf("Body:\n")
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		fmt.Printf("%s\n", body)

		// Respond to the request
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Request received"))

		fmt.Printf("--- Done\n")
	})

	fmt.Printf("Listening to http://localhost%s...\n", port)

	srv := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return 1
	}

	return 0
}
