package main

import (
	"github.com/Mire0726/Genkiyoho/backend/server"
	"flag"
)

func main() {
    	// コマンドライン引数の解析
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "Address to listen on")
	flag.Parse()

	// サーバーの起動
	server.Serve(addr)
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