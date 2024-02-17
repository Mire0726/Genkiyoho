package model

import (
    "database/sql"
    "log"
    "time"
)

// User 構造体の定義（重複した定義を削除）
type User struct {
    ID        int       `json:"id"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// GetAllUsers データベースから全てのユーザーを取得する
func GetAllUsers(db *sql.DB) ([]User, error) {
    rows, err := db.Query("SELECT id, email, password, name, created_at, updated_at FROM users")
    if err != nil {
        log.Printf("Error querying users from database: %v", err)
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        var createdAtStr, updatedAtStr string // 中間変数を用意

        if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &createdAtStr, &updatedAtStr); err != nil {
            log.Printf("Error scanning user: %v", err)
            return nil, err
        }

        layout := "2006-01-02 15:04:05"
        // createdAt と updatedAt を time.Time 型に変換
        user.CreatedAt, err = time.Parse(layout, createdAtStr)
        if err != nil {
            log.Printf("Error parsing createdAt for user %v: %v", user.ID, err)
            return nil, err
        }

        user.UpdatedAt, err = time.Parse(layout, updatedAtStr)
        if err != nil {
            log.Printf("Error parsing updatedAt for user %v: %v", user.ID, err)
            return nil, err
        }

        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Error during rows iteration: %v", err)
        return nil, err
    }

    return users, nil
}
