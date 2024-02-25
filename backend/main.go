package main

import (
	"flag"

	"github.com/Mire0726/Genkiyoho/backend/server"
	
)

func main() {
    var addr string
    flag.StringVar(&addr, "addr", ":8080", "Address to listen on")
    flag.Parse()

    // サーバーの起動
    server.Serve(addr)
}



// package main

// import (
// 	"github.com/Mire0726/Genkiyoho/backend/server"
// 	"flag"
// 	"log"
// 	"github.com/Mire0726/Genkiyoho/backend/server/db"
// )

// func main() {
//     // データベース接続の取得
//     conn, err := db.ConnectToDB()
//     if err != nil {
//         log.Fatalf("Failed to connect to database: %v", err)
//     }
//     defer conn.Close() // 関数終了時にデータベース接続を閉じる

// 	var addr string
// 	flag.StringVar(&addr, "addr", ":8080", "Address to listen on")
// 	flag.Parse()

// 	// サーバーの起動
// 	server.Serve(addr,conn)
// }


