package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
    connStr := "user=postgres password=yourpassword dbname=users sslmode=disable"
    return sql.Open("postgres", connStr)
}