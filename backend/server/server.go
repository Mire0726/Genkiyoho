package server

import (
	"log"
	"net/http"

	"github.com/Mire0726/Genkiyoho/backend/server/handler"
	// "github.com/Mire0726/Genkiyoho/backend/server/http/middleware"
	_ "github.com/go-sql-driver/mysql" // MySQLドライバーをインポート
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

// Serve はHTTPサーバを起動します。データベース接続を引数に追加。
func Serve(addr string) {
    e := echo.New()

    
// ミドルウェアの設定
    // panicが発生した場合の処理
	e.Use(echomiddleware.Recover())
	// CORSの設定
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		Skipper:      echomiddleware.DefaultCORSConfig.Skipper,
		AllowOrigins: echomiddleware.DefaultCORSConfig.AllowOrigins,
		AllowMethods: echomiddleware.DefaultCORSConfig.AllowMethods,
		AllowHeaders: []string{"Content-Type,Accept,Origin,x-token"},
	}))




    // ルーティングの設定
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Welcome to Genkiyoho!")
    })
    e.POST("/user/create", handler.HandleUserCreate())
    e.GET("/users", handler.HandleGetUser()) 

    // authAPI := e.Group("", middleware.AuthenticateMiddleware())

    

    /* ===== サーバの起動 ===== */

    log.Printf("Server running on %s", addr)
    if err := e.Start(addr); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

