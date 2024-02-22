package main

import (
	"github.com/Mire0726/Genkiyoho/backend/server"
	"flag"
	"log"
	"github.com/Mire0726/Genkiyoho/backend/server/db"
)

func main() {
    // データベース接続の取得
    conn, err := db.ConnectToDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer conn.Close() // 関数終了時にデータベース接続を閉じる

	var addr string
	flag.StringVar(&addr, "addr", ":8080", "Address to listen on")
	flag.Parse()

	// サーバーの起動
	server.Serve(addr,conn)
}



// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	//Ginフレームワークのデフォルトの設定を使用してルータを作成
// 	router := gin.Default()
	
// 	// ルートハンドラの定義
// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Hello, World!",
// 		})
// 	})

// 	// サーバー起動
// 	router.Run(":8080")
// }