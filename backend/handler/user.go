package handler

import (
    "github.com/Mire0726/Genkiyoho/backend/db"
    "github.com/Mire0726/Genkiyoho/backend/model"
    
    "encoding/json"
    "net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    dbConn, err := db.ConnectToDB()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer dbConn.Close()

    rows, err := dbConn.Query("SELECT id, name FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    users := make([]model.User, 0)
    for rows.Next() {
        var u model.User
        if err := rows.Scan(&u.ID, &u.Name); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, u)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}