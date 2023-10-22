package handlers

import (
    "fmt"
    "net/http"
)

func FooBar(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "bar")
}
