package main

import (
	"fmt"
	"gRMS-client/client"
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
	err := c.Connect("localhost:8080", "ws")
	if err != nil {
		log.Fatalln("error connecting: ", err)
	}

	d := data.NewDataHandler(c)

	l := logger.NewChatLogger(d)
	l.StartLogging()

	uphandler := updates.NewUpdatesHandler(c.GetUpdatesChan(), l, d)
	uphandler.Start()
}
