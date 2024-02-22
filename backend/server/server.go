package server

import (
	"github.com/Mire0726/Genkiyoho/backend/server/handler"
	"net/http"
	"github.com/labstack/echo/v4"
	"log"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"database/sql"
)

// Serve はHTTPサーバを起動します。データベース接続を引数に追加。
func Serve(addr string, db *sql.DB) {
    e := echo.New()

    // ミドルウェアの設定
    e.Use(echomiddleware.Recover())
    e.Use(echomiddleware.CORS())

    // ルーティングの設定
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Welcome to Genkiyoho!")
    })

    // AppHandlerのインスタンスを作成し、データベース接続を渡す
    appHandler := &handler.AppHandler{DB: db}

    // ルーティングの設定
    e.GET("/users", appHandler.HandleGetUser) // 修正: appHandlerを使用
    e.POST("/user/create", appHandler.HandleUserCreate)

    // サーバーの起動
    log.Printf("Server running on %s", addr)
    if err := e.Start(addr); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

