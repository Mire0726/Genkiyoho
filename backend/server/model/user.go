package model

import (
	"database/sql"
	"log"
	"time"
	// "github.com/Mire0726/Genkiyoho/backend/server/db"
)

// User 構造体の定義（重複した定義を削除）
type User struct {
    ID        int       `json:"id"`
    AuthToken string    `json:"authtoken"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

var db *sql.DB 


// InsertUser データベースにユーザレコードを登録する
func InsertUser(db *sql.DB, record *User) error {
    _, err := db.Exec(
        "INSERT INTO users (auth_token, email, password, name) VALUES (?, ?, ?, ?)",
        record.AuthToken,
        record.Email,
        record.Password,
        record.Name,
    )
    if err != nil {
        log.Printf("ユーザーの登録に失敗しました: %v", err)
        return err
    }
    log.Println("ユーザーが正常に登録されました。")
    return nil
}


// SelectUserByAuthToken 認証トークンに紐づくユーザ情報をデータベースから取得する
func SelectUserByAuthToken(db *sql.DB, token string) (*User, error) {
    var user User
    query := `SELECT id, auth_token, email, password, name, created_at, updated_at FROM users WHERE auth_token = ?`
    err := db.QueryRow(query, token).Scan(&user.ID, &user.AuthToken, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // ユーザが見つからない場合
        }
        log.Printf("Failed to find user by auth token: %v", err)
        return nil, err
    }

    return &user, nil
}


// GetAllUsers データベースから全ユーザを取得する
func GetAllUsers() ([]User, error) {
    rows, err := db.Query("SELECT id, auth_token, email, password, name, created_at, updated_at FROM users")
    if err != nil {
        log.Printf("Error querying users from database: %v", err)
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User

        if err := rows.Scan(&user.ID, &user.AuthToken, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
            log.Printf("Error scanning user: %v", err)
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