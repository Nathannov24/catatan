package main

import (
	"transaksi/config"
	m "transaksi/middlewares"
	"transaksi/routers"
)

func main() {
	config.InitDB()
	e := routers.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
