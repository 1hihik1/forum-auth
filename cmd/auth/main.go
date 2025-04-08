package main

import "github.com/DrusGalkin/Auth-gRPC/internal/app"

func main() {
	go app.Run()
	go app.StartGRPCServer()
	select {}
}
