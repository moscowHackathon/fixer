package callback

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/moscowHackathon/fixer/service"
	"github.com/moscowHackathon/fixer/slackrequest"
	"github.com/moscowHackathon/slack"
)

var (
	// API slack web API
	API *slack.Client
)

// HandleHome handle slack requests
func HandleHome(w http.ResponseWriter, r *http.Request) {
	payload := r.PostFormValue("payload")
	data := &slackrequest.Payload{}
	err := json.Unmarshal([]byte(payload), data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't unmarshal json: %s", err)
	}
	answer := data.Actions[0].Value
	response, err := service.Answer(data.Channel.ID, answer)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	if strings.Contains(response.Message, "complete") == true {
		w.Write([]byte(response.Message))
		return
	}
	response, err = service.Question(data.Channel.ID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	message := data.OriginalMessage
	message.Text = "So, " + data.Actions[0].Name
	message.Attachments[0].CallbackID = response.ID
	message.Attachments[0].Text = response.Message
	if API == nil {
		return
	}
	_, _, _, err = API.UpdateMessage(data.Channel.ID, data.MessageTS, message.Text, message.Attachments)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
}

// HandleStart handle slack slash command /fix requests
func HandleStart(w http.ResponseWriter, r *http.Request) {
	userID := r.PostFormValue("user_id")
	userName := r.PostFormValue("user_name")
	_, _, channelID, err := API.OpenIMChannel(userID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	API.InviteUserToChannel(channelID, userID)
	response, err := service.Start(channelID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	response, err = service.Question(channelID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	params := slack.NewPostMessageParameters()
	params.Attachments = []slack.Attachment{service.GenerateMessageForSlack2(response.Message, response.ID)}
	_, _, err = API.PostMessage(channelID, "Hello, "+userName+"!", params)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
}

// Run start listening for slack requests
func Run(cert, key string) {

	http.HandleFunc("/", HandleHome)
	http.HandleFunc("/start", HandleStart)
	http.Handle("/favicon.ico", http.FileServer(http.Dir("/var/www/html")))

	http.ListenAndServeTLS(
		":443",
		cert,
		key,
		//"/etc/letsencrypt/live/arkon-bot.ru/fullchain.pem",
		//"/etc/letsencrypt/live/arkon-bot.ru/privkey.pem",
		nil,
	)
}
