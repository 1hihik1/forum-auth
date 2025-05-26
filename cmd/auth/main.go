package main

import "github.com/1hihik1/forum-auth/internal/app"

func main() {
	app.LoggerRun()
	go app.Run()
	go app.StartGRPCServer()
	select {}
	//password, err := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(password))
}
