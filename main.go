package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/101nl/slack"
	"github.com/kr/pretty"
	"github.com/moscowHackathon/fixer/service"
)

var channelID string
var response service.GetResponse

func main() {
	token := flag.String("slack-token", "", "Token from slack")
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
				if strings.HasPrefix(ev.Channel, "D") {
					if ev.BotID != "" {
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
					fmt.Println(" >>> ============================== Start ")
					fmt.Println(response)
					fmt.Println("<<<  ============================== Start ")


					//TODO: А вот тут уже отрисуем кнопочки? или не тут?!
					atachment2 := service.GenerateMessageForSlack2("А вот тут будет текст вопроса",)
					msg := slack.NewPostMessageParameters()
					msg.Attachments = []slack.Attachment{atachment2}
					rtm.PostMessage(ev.Channel /*"C1NBBSKEE"*/, "Тут будет заголовок окна", msg)

/*
				params := slack.NewPostMessageParameters()
				attachment := service.GenerateMessageForSlack("qwe")
				params.Attachments = []slack.Attachment{attachment}
				//responseChannel, responseTime, err := rtm.PostMessage(channelID, "Текст !!!", params)
				rtm.PostMessage(channelID, "Текст !!!", params)
				//chan id = C1NBBSKEE

 */
				}

				//=======================================================================================================
				//				rtm.SendMessage(rtm.NewOutgoingMessage("Сам дурак. Ответ эксперта - " + strconv.Itoa(int(response.ID)), channelID))
				rtm.SendMessage(rtm.NewOutgoingMessage("Сам дурак. Ответ эксперта - "+response.ID, channelID))

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

