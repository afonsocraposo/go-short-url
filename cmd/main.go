// cmd/main.go
package main

import (
    "fmt"
    "net/http"
    "shorturl/internal/handlers"
)

func main() {
    http.HandleFunc("/foo", handlers.FooBar)
    http.HandleFunc("/", handlers.ShortUrlHandler)
    fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}