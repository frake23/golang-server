package main

import (
	"./server"
	"fmt"
)

const (
	fileDB string = "./repository/users.json"
)

func main() {
	app := server.InitApp(fileDB)

	port := "8080"
	if err := app.Run(port); err != nil {
		fmt.Printf("%s", err.Error())
		// log.Fatalf("%s", err.Error())
	}
}
