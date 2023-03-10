package command

import (
	"bufio"
	"fmt"
	"gRMS-client/client"
	"gRMS-client/data"
	logger "gRMS-client/log"
	"os"
	"strings"

	"github.com/gookit/color"
)

func Listen(c client.Client, d data.DataHandler) {
	var reader = bufio.NewReader(os.Stdin)
	yellow := color.FgYellow
	blue := color.FgBlue

	for {
		var command string
		fmt.Printf("%s> ", logger.Prompt)
		fmt.Scan(&command)

		if command == "help" {
			helpline := strings.Builder{}

			helpline.WriteString(yellow.Sprintf("~> chat_in <chat_id> (go into chat to send messages)\n"))
			helpline.WriteString(yellow.Sprintf("~> list_chats (list all available chats)\n"))
			helpline.WriteString(yellow.Sprintf("~> new_chat (create a new chat)\n"))
			helpline.WriteString(yellow.Sprintf("~> help (displays this help message)\n"))
			helpline.WriteString(yellow.Sprintf("~> exit (close the application)\n"))

			fmt.Print(helpline.String())
		}

		if command == "list_chats" {
			if self := d.GetSelf(); self != nil {
				for _, v := range d.GetSelf().Chats {
					if c := d.GetChat(v.ChatID); c != nil {
						blue.Light().Printf("  - %s ", c.Title)
						fmt.Printf("(id:%d)\n", c.ID)
					}
				}
			} else {
				fmt.Println("data not found")
			}
		}

		if command == "chat_in" {
			var chatID uint64
			fmt.Scanln(&chatID)

			if chat := d.GetChat(chatID); chat != nil {
				logger.Prompt = blue.Sprint("(" + chat.Title + ")")
				for {
					fmt.Printf("(%s)> ", chat.Title)
					text := GetString(reader)
					if text == "" {
						continue
					}

					if text == ">back" {
						break
					}

					c.SendMessage(chatID, strings.TrimSpace(text), 0)
				}
			} else {
				fmt.Println("chat not found")
			}
		} else if command == "new_chat" {
			blue.Light().Print("Title: ")
			title := GetString(reader)
			fmt.Print("enter the usernames to add: ")
			usernames := GetString(reader)
			usernames = strings.ReplaceAll(usernames, "\n", "")

			c.CreateChat(title, strings.Split(usernames, " "))
		} else if command == "exit" {
			c.Close()
		}

		logger.Prompt = "~"
	}
}

func GetString(in *bufio.Reader) string {
	s, _ := in.ReadString('\n')
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.TrimSpace(s)
	return s
}
