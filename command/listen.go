package command

import (
	"bufio"
	"fmt"
	"gRMS-client/client"
	"gRMS-client/data"
	logger "gRMS-client/log"
	"log"
	"os"
	"strconv"
	"strings"
)

func Listen(c client.Client, d data.DataHandler) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s> ", logger.Prompt)
		in, err := reader.ReadString('\n')
		in = strings.TrimSpace(in)
		in = strings.ReplaceAll(in, "\r", "")
		in = strings.ReplaceAll(in, "\n", "")

		if err != nil {
			log.Fatalln("error while reding input", err)
		}

		if in == "list_chats" {
			if self := d.GetSelf(); self != nil {
				for _, v := range d.GetSelf().Chats {
					if c := d.GetChat(v.ChatID); c != nil {
						fmt.Printf("%s (id:%d)\n", c.Title, c.ID)
					}
				}
			} else {
				fmt.Println("data not found")
			}
		}

		if s := strings.Split(in, " "); s[0] == "chat_in" {
			chatID, err := strconv.ParseUint(s[1], 10, 64)
			if err != nil {
				fmt.Println("invalid chat id")
				continue
			}

			if chat := d.GetChat(chatID); chat != nil {
				logger.Prompt = "(" + chat.Title + ")"
				for {
					fmt.Printf("(%s)> ", chat.Title)
					in, err := reader.ReadString('\n')
					in = strings.TrimSpace(in)
					in = strings.ReplaceAll(in, "\r", "")
					in = strings.ReplaceAll(in, "\n", "")

					if err != nil {
						log.Fatalln("error while reading input", err)
					}
					if in == ">back" {
						break
					}
					c.SendMessage(chatID, strings.TrimSpace(in), 0)
				}
			} else {
				fmt.Println("chat not found")
			}
		} else if s[0] == "new_chat" {
			title := strings.Join(s[1:], " ")
			fmt.Print("enter the usernames to add: ")
			var usernames string
			fmt.Scanln(&usernames)

			c.CreateChat(title, strings.Split(usernames, " "))
		}

		logger.Prompt = "~"
	}
}
