package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kr/pretty"
	"github.com/moscowHackathon/fixer/callback"
	"github.com/moscowHackathon/fixer/service"
	"github.com/moscowHackathon/slack"
)

var channelID string
var response service.GetResponse
var textDebug = true

func main() {
	token := flag.String("slack-token", "", "Token from slack")
	cert := flag.String("cert", "", "Path to SSL cert file")
	key := flag.String("key", "", "Path to SSL key file")
	flag.Parse()

	if token == nil || *token == "" {
		flag.Usage()
		os.Exit(1)
	}

	api := slack.New(*token)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	// Find the user to post as.
	authTest, err := api.AuthTest()
	if err != nil {
		fmt.Printf("Error getting channels: %s\n", err)
		return
	}

	botID := authTest.UserID

	pretty.Println("botID", botID)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	if cert != nil && *cert != "" && key != nil && *key != "" {
		callback.API = api
		go callback.Run(*cert, *key)
	}

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
			// Ignore hello

			case *slack.ConnectedEvent:
				// Replace #general with your Channel ID
				//rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "C1NBBSKEE"))

			case *slack.MessageEvent:
				fmt.Printf("Message: %+v\n", ev)
				//Проверка на приват
				if strings.HasPrefix(ev.Channel, "D") {
					if ev.BotID != "" || ev.SubType == "bot_message" || ev.SubType == "message_changed" {
						break
					}
					if len(ev.Attachments) == 1 && ev.Attachments[0].CallbackID != "" {
						break
					}

					if textDebug {

						if strings.Contains(ev.Text, "complete") == true {
							rtm.SendMessage(rtm.NewOutgoingMessage("Компонент определен. Это - '"+ev.Text+"'", channelID))
							break
						}

						response, _ = service.Answer(channelID, ev.Text)
						rtm.SendMessage(rtm.NewOutgoingMessage("Ты ответил: "+ev.Text, channelID))

						response, _ = service.Question(channelID)
						rtm.SendMessage(rtm.NewOutgoingMessage("Вот тебе вопрос: '"+response.Message+"'. Варианты ответов : 1,0,-1", channelID))

						break
					}

				}

				//Игнорим сообщения, не предназначеные боту
				if strings.Contains(ev.Text, "<@"+botID+">") == false {
					break
				}

				// Если мы не в приватном канале - послать приглашение
				if channelID != ev.Channel {
					_, _, channelID, err = rtm.OpenIMChannel(ev.User)
					rtm.InviteUserToChannel(channelID, ev.User)
					response, _ = service.Start(channelID)

					if textDebug {
						response, _ = service.Question(channelID)
						rtm.SendMessage(rtm.NewOutgoingMessage("Вот тебе вопрос: '"+response.Message+"'. Варианты ответов : 1,0,-1", channelID))
						break
					}

					//Рисуем кнопочки в привате
					atachment2 := service.GenerateMessageForSlack2("А вот тут будет текст вопроса", "id")
					msg := slack.NewPostMessageParameters()
					msg.Attachments = []slack.Attachment{atachment2}
					rtm.SendMessage(rtm.NewOutgoingMessage("Привет пользователь - "+response.ID, channelID))
					rtm.PostMessage(channelID /*"C1NBBSKEE"*/, "Тут будет заголовок окна", msg)
				}

				//=======================================================================================================
				//rtm.SendMessage(rtm.NewOutgoingMessage("Сам дурак. Ответ эксперта - " + strconv.Itoa(int(response.ID)), channelID))
				//rtm.SendMessage(rtm.NewOutgoingMessage("Сам дурак. Ответ эксперта - "+response.ID, channelID))

			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
				fmt.Printf("Current latency: %v\n", ev.Value)

			case *slack.RTMError:
				fmt.Printf("\033[0;31mError:\033[0m %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("\033[0;31mInvalid credentials\033[0m")
				os.Exit(2)

			default:
				// Ignore other events..
				pretty.Println("Unexpected:", msg.Data)
			}
		}
	}
}
