package handlers

import (
    "fmt"
    "net/http"
    "net/url"
    "github.com/afonsocraposo/go-short-url/internal/helpers"
    "strings"
)

const Size int = 16

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getUrl(w, r)
    case http.MethodPost:
        generateShortUrl(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getUrl(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    parts := strings.Split(path, "/")
    if len(parts) == 2 {
        id := parts[1]
        if len(id) != Size {
            http.Error(w, "Invalid URL format", http.StatusBadRequest)
            return
        }
        url, err := helpers.GetVar(id)
        if url == "" {
            if err == nil {
                http.Error(w, "URL not found", http.StatusNotFound)
            } else {
                http.Error(w, "Something went wrong", http.StatusInternalServerError)
            }
            return
        }
        http.Redirect(w, r, url, 301)
    } else {
        http.Error(w, "Invalid URL format", http.StatusBadRequest)
    }
}

func isUrl(str string) bool {
    u, err := url.Parse(str)
    return err == nil && u.Scheme != "" && u.Host != ""
}

func generateShortUrl(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }

    url := r.FormValue("url")

    if !isUrl(url) {
        http.Error(w, "Invalid URL format", http.StatusBadRequest)
        return
    }

    hash := helpers.HashString(url, Size)

    err = helpers.SetVar(hash, url)
    if err != nil {
        http.Error(w, "Something went wrong", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "http://localhost:8080/%s", hash)
}
