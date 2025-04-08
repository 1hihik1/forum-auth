package main

import "github.com/DrusGalkin/forum-auth-grpc/internal/app"

func main() {
	go app.Run()
	go app.StartGRPCServer()
	select {}
}
