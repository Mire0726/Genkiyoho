package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "fmt"
)


func ConnectToDB() (*sql.DB, error) {
	connStr := "host=localhost port=5433 user=postgres password=postgres dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// 接続の開始に失敗した場合のエラーログ
		log.Printf("データベースへの接続の開始に失敗しました: %v", err)
		return nil, fmt.Errorf("データベースへの接続の開始に失敗: %w", err)
	}

	// データベース接続の確認
	if err = db.Ping(); err != nil {
		// 接続の確認に失敗した場合のエラーログ
		log.Printf("データベースへの接続の確認に失敗しました: %v", err)
		return nil, fmt.Errorf("データベースへの接続の確認に失敗: %w", err)
	}

	log.Println("データベースへの接続に成功しました。")
	return db, nil
}
