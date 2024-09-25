package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/skip2/go-qrcode"
)

func main() {
	// Handle requests to /qr
	http.HandleFunc("/qr", func(w http.ResponseWriter, r *http.Request) {
		// Get the 'url' query parameter
		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
			return
		}
		size := r.URL.Query().Get("size")
		if size == "" {
			http.Error(w, "Missing 'size' query parameter", http.StatusBadRequest)
			return
		}
		sizeInt, err := strconv.Atoi(size)
		if err != nil {
			http.Error(w, "Invalid 'size' query parameter", http.StatusBadRequest)
			return
		}

		// Generate the QR code
		qr, err := qrcode.New(url, qrcode.Highest)
		if err != nil {
			http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
			return
		}

		// Serve the QR code as a PNG image
		w.Header().Set("Content-Type", "image/png")
		err = qr.Write(sizeInt, w)
		if err != nil {
			http.Error(w, "Failed to write image", http.StatusInternalServerError)
		}
	})

	// Start the server
	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
