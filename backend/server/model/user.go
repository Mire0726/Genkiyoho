package model

import (
	"database/sql"
	"log"
	"time"
	"github.com/Mire0726/Genkiyoho/backend/server/db"
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

// InsertUser データベースにユーザレコードを登録する
func InsertUser(record *User) error  {
    log.Println("model,line24")
    _, err := db.Conn.Exec(
        "INSERT INTO users (auth_token, email, password, name) VALUES (?, ?, ?, ?)",
        record.AuthToken,
        record.Email,
        record.Password,
        record.Name,
    )
    if err != nil {
        log.Printf("Error inserting user into database: %v", err) // ログ追加
        return err
    }
    log.Println("User successfully registered.") // 成功メッセージもログに記録
    return nil
}

// SelectUserByAuthToken auth_tokenを条件にレコードを取得する
func SelectUserByAuthToken(authToken string) (*User, error) {
    user := &User{}
    err := db.Conn.QueryRow("SELECT id, auth_token FROM `users` WHERE auth_token=?", authToken).Scan(&user.ID, &user.AuthToken)
    if err != nil {
        if err == sql.ErrNoRows {
            // レコードが見つからない場合はnilを返す
            return nil, nil
        }
        // その他のエラー
        log.Printf("Error querying user by auth token: %v", err)
        return nil, err
    }

    return user, nil
}
// SelectUserByPrimaryKey 主キーを条件にレコードを取得する
func SelectUserByPrimaryKey(userID int) (*User, error) {
	row := db.Conn.QueryRow("SELECT * FROM user WHERE id=?", userID)
	return convertToUser(row)
}


// UpdateUserByPrimaryKey 主キーを条件にレコードを更新する
func UpdateUserByPrimaryKey(record *User) error {
	if _, err := db.Conn.Exec(
		"UPDATE user SET name=? WHERE id=?",
		record.Name,
		record.ID,
	); err != nil {
		return err
	}
	return nil
}
// GetAllUsers データベースから全ユーザを取得する
func GetAllUsers() ([]User, error) {
    rows, err := db.Conn.Query("SELECT id, auth_token, email, password, name, created_at, updated_at FROM users")
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

// convertToUser rowデータをUserデータへ変換する
func convertToUser(row *sql.Row) (*User, error) {
	var user User
	if err := row.Scan(&user.ID, &user.AuthToken, &user.Name, &user.Email, &user.Password,&user.CreatedAt,&user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}