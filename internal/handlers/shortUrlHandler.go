package handlers

import (
    "os"
    "errors"
    "fmt"
    "net/http"
    "net/url"
    "github.com/afonsocraposo/go-short-url/internal/helpers"
    "strings"
)

const Size int = 5

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getHandler(w, r)
    case http.MethodPost:
        generateShortUrl(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    fp := "./internal/static" + path

    parts := strings.Split(path, "/")

    switch len(parts) {
    case 2:
        fallthrough
    case 3:
        _, err := os.Stat(fp)
        if errors.Is(err, os.ErrNotExist) {
            id := parts[1]
            getUrl(w, r, id)
            return
        }
        fallthrough
    default:
        http.ServeFile(w, r, fp)
    }
}

func getUrl(w http.ResponseWriter, r *http.Request, id string) {
    if len(id) != Size {
        http.Redirect(w, r, "/", 301)
        return
    }
    url, _ := helpers.GetVar(id)
    if url == "" {
        http.Redirect(w, r, "/", 301)
        return
    }
    http.Redirect(w, r, url, 301)
}

func isUrlValid(str string) bool {
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

    if !isUrlValid(url) {
        http.Error(w, "Invalid URL format", http.StatusBadRequest)
        return
    }

    hash := helpers.HashString(url, Size)

    err = helpers.SetVar(hash, url)
    if err != nil {
        http.Error(w, "Something went wrong", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "http://127.0.0.1:8080/%s", hash)
}
