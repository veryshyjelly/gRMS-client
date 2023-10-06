package command

import (
	"bufio"
	"fmt"
	"gRMS-client/client"
	"gRMS-client/data"
	logger "gRMS-client/log"
	"github.com/manifoldco/promptui"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

func Listen(c client.Client, d data.Handler) {
	var reader = bufio.NewReader(os.Stdin)
	yellow := color.FgYellow
	blue := color.FgBlue
	cyan := color.FgCyan
	red := color.FgRed
	green := color.FgGreen

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
			CHATLOOP:
				for {
					fmt.Printf("%s> ", logger.Prompt)
					text := GetString(reader)
					if text == "" {
						continue
					}

					if text[0] == '>' {
						if command := strings.Split(text[1:], " "); command[0] == "back" {
							break
						} else if command[0] == "add" {
							if len(command) == 1 {
								red.Printf("no username given -_-\n")
								continue
							}

							c.AddToChat(chatID, command[1:])
						} else if command[0] == "history" {
							mess := d.GetMessages(chatID)

							var last int
							if len(command) != 1 {
								last, _ = strconv.Atoi(command[1])
								if last > len(mess) {
									last = 0
								} else {
									last = len(mess) - last
								}
							}

							for _, m := range mess[last:] {
								m.Log(chat, d.GetUser(m.From))
							}

						} else if command[0] == "" {
							fmt.Print("\033[A\033[K")

							prompt := promptui.Select{
								Label:    "Options",
								HideHelp: true,
								Items: []string{"Change Title", "Add Participants",
									"Remove Participant", "Invite via link", "Promote to Admin", "Dismiss as Admin", "Go to Home"},
							}
							x, _, err := prompt.Run()
							if err != nil {
								red.Printf("error while selecting %v", err)
							}

							switch x {
							case 0:
								blue.Light().Print("Title: ")
								var title string
								for title == "" {
									title = GetString(reader)
								}
								c.UpdateChatTitle(chatID, title)
							case 1:
								cyan.Print("Enter usernames to add: ")
								usernames := GetString(reader)
								c.AddToChat(chatID, strings.Split(usernames, " "))
							case 2:
								cyan.Print("Enter username to kick: ")
								usernames := GetString(reader)
								c.RemoveFromChat(chatID, strings.Split(usernames, " "))
							case 3:
								break
							case 4:
								cyan.Print("Enter username to promote: ")
								usernames := GetString(reader)
								c.PromoteUsers(chatID, strings.Split(usernames, " "))
							case 5:
								cyan.Print("Enter username to demote: ")
								usernames := GetString(reader)
								c.DemoteUsers(chatID, strings.Split(usernames, " "))
							default:
								break CHATLOOP
							}

							//fmt.Println(x)
						} else if command[0] == "members" {
							cyan.Println("Members:")
							for _, v := range chat.Members {
								green.Light().Printf("  - %v\n", d.GetUser(v.UserID).Username)
							}
						} else if command[0] == "admins" {
							cyan.Println("Admins:")
							for _, v := range chat.Admins {
								green.Light().Printf("  - %v\n", d.GetUser(v.UserID).Username)
							}
						} else if command[0] == "leave" {
							prompt := promptui.Select{Label: "Confirm exit?",
								HideHelp: true, Items: []string{"YES", "NO"},
							}
							_, result, err := prompt.Run()
							if err != nil {
								red.Printf("error while selecting %v", err)
							}

							if result == "YES" {
								c.LeaveChat(chatID)
								break
							}
						} else if command[0] == "rf" {
							c.GetChat(chatID)
							chat = d.GetChat(chatID)
							logger.Prompt = chat.Title
						} else if command[0] == "help" {
							helpline := strings.Builder{}

							helpline.WriteString(yellow.Sprintf(">back (go back from this chat)\n"))
							helpline.WriteString(yellow.Sprintf(">add <list of usernames> (add users to the chat)\n"))
							helpline.WriteString(yellow.Sprintf(">history <amount> (show history)\n"))
							helpline.WriteString(yellow.Sprintf(">help (displays this message)\n"))
							helpline.WriteString(yellow.Sprintf("> (chat settings)"))

							fmt.Print(helpline.String())
						}

					} else {
						c.SendMessage(chatID, strings.TrimSpace(text), 0)
					}
				}
			} else {
				fmt.Println("chat not found")
			}
		} else if command == "new_chat" {
			blue.Light().Print("Title: ")
			var title string
			for title == "" {
				title = GetString(reader)
			}
			cyan.Print("enter the usernames to add: ")
			usernames := GetString(reader)
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