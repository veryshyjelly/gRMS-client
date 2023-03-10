package main

import (
	"fmt"
	"gRMS-client/client"
	"gRMS-client/command"
	"gRMS-client/data"
	logger "gRMS-client/log"
	"gRMS-client/updates"
	"log"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {
	var host string
	fmt.Print("Enter host address: ")
	fmt.Scanln(&host)

	if host == "" {
		host = "localhost:8080"
	}

	fmt.Print("\033[A\033[K")

	prompt := promptui.Select{
		Label: "Sign In or Sign Up?",
		Items: []string{"Sign In", "Sign Up"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalln("error reading prompt", err)
	}

	var username, password string
	if result == "Sign In" {
		fmt.Print("username: ")
		fmt.Scanln(&username)
		fmt.Print("\033[A\033[Kpassword: ")
		fmt.Scanln(&password)
		fmt.Print("\033[A\033[K")
	} else {
		log.Fatalln("Sign Up option is not build yet.")
	}

	c := client.NewClient(username, password)
	d := data.NewDataHandler(c)

	l := logger.NewChatLogger(d)
	go l.StartLogging()

	uphandler := updates.NewUpdatesHandler(c.GetUpdatesChan(), l, d)
	go uphandler.Start()

	go command.Listen(c, d)

	err = c.Connect(host, "ws")
	if err != nil {
		log.Fatalln("error connecting: ", err)
	}

	time.Sleep(time.Millisecond * 100)
}
