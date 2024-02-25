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


