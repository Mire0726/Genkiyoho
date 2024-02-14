package server

import (
    "github.com/Mire0726/Genkiyoho/backend/handler"
    "net/http"
)

func StartServer() {
    http.HandleFunc("/users", handler.GetUser)
    http.ListenAndServe(":8080", nil)
}

