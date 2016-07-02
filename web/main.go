package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"strconv"

	"github.com/moscowHackathon/fixer/slackrequest"
)

var responses = map[string]uint64{}

func handler(w http.ResponseWriter, r *http.Request) {
	payload := r.PostFormValue("payload")
	data := &slackrequest.Payload{}
	err := json.Unmarshal([]byte(payload), data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't unmarshal json: %s", err)
	}
	responses[data.User.ID]++
	message := data.OriginalMessage
	message.Text = "Question #" + strconv.FormatUint(responses[data.User.ID], 10)
	message.Attachments[0].CallbackID = message.Text
	response, err := json.Marshal(message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(response)
}

func main() {
	cert := flag.String("cert", "", "Path to SSL cert file")
	key := flag.String("key", "", "Path to SSL key file")
	flag.Parse()

	if cert == nil || *cert == "" || key == nil || *key == "" {
		flag.Usage()
		os.Exit(1)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(
		":443",
		*cert,
		*key,
		//"/etc/letsencrypt/live/arkon-bot.ru/fullchain.pem",
		//"/etc/letsencrypt/live/arkon-bot.ru/privkey.pem",
		nil,
	)
}
