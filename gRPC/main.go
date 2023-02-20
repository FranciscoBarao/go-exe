package main

import (
	"bookshop/client"
	"bookshop/server"
	"time"
)

func main() {
	go server.NewServer()
	time.Sleep(5 * time.Second)
	client.NewClient()
	time.Sleep(100 * time.Second)
}
