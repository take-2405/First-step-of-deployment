package main

import (
	"flag"
	"F-deployment/pkg/server"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	//port := os.Getenv("PORT")
	//if port == "" {
	//	log.Fatal("$PORT must be set")
	//}
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	server.Server.Run(addr)
}
