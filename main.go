package main

import (
	"fmt"
	"gRMS-client/client"
	"gRMS-client/command"
	"gRMS-client/data"
	logger "gRMS-client/log"
	"gRMS-client/updates"
	"log"
)

func main() {
	var username, password string
	fmt.Print("username: ")
	fmt.Scanln(&username)
	fmt.Print("password: ")
	fmt.Scanln(&password)

	c := client.NewClient(username, password)
	d := data.NewDataHandler(c)

	l := logger.NewChatLogger(d)
	go l.StartLogging()

	uphandler := updates.NewUpdatesHandler(c.GetUpdatesChan(), l, d)
	go uphandler.Start()

	go command.Listen(c, d)

	err := c.Connect("localhost:8080", "ws")
	if err != nil {
		log.Fatalln("error connecting: ", err)
	}
}
