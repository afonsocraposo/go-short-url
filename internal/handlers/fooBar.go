package handlers

import (
    "fmt"
    "net/http"
)

func FooBar(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "bar")
}
