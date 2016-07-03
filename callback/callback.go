package callback

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/kr/pretty"
	"github.com/moscowHackathon/fixer/slackrequest"
	"github.com/moscowHackathon/slack"
)

var (
	responses = map[string]uint64{}

	// API slack web API
	API *slack.Client
)

// Handle handle slack requests
func Handle(w http.ResponseWriter, r *http.Request) {
	payload := r.PostFormValue("payload")
	data := &slackrequest.Payload{}
	err := json.Unmarshal([]byte(payload), data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't unmarshal json: %s", err)
	}
	responses[data.User.ID]++
	message := data.OriginalMessage
	message.Text = "So, " + data.Actions[0].Name + ". The next question is #" + strconv.FormatUint(responses[data.User.ID], 10)
	message.Attachments[0].CallbackID = message.Text
	message.BotID = "asdasd"
	if API == nil {
		return
	}
	a, b, c, d := API.UpdateMessage(data.Channel.ID, data.MessageTS, message.Text, message.Attachments)
	pretty.Println(a, b, c, d)
}

// Run start listening for slack requests
func Run(cert, key string) {
	http.HandleFunc("/", Handle)
	http.ListenAndServeTLS(
		":443",
		cert,
		key,
		//"/etc/letsencrypt/live/arkon-bot.ru/fullchain.pem",
		//"/etc/letsencrypt/live/arkon-bot.ru/privkey.pem",
		nil,
	)
}
