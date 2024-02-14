package server

import (
    //インポートエラー
    "../handler"
    "net/http"
)

func StartServer() {
    http.HandleFunc("/users", handler.GetUser)
    http.ListenAndServe(":8080", nil)
}

