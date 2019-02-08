package main

import (
    "github.com/yosssi/ace"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    tpl, err := ace.Load("index", "", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tpl.Execute(w, map[string]string{"Title": "hello", "H1": "hello world"}); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
