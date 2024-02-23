package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // MySQLドライバーをインポート
    "log"
    "fmt"
)

func ConnectToDB() (*sql.DB, error) {
    connStr := "root:mysql@tcp(localhost:3307)/db?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := sql.Open("mysql", connStr)
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
