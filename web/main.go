package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"github.com/moscowHackathon/fixer/service"
)

func handler(w http.ResponseWriter, r *http.Request) {
	buf, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprintf(os.Stdout, "ERROR: %s", err)
		return
	}
	params, err := url.ParseQuery(string(buf))
	if err != nil {
		fmt.Fprintf(os.Stdout, "ERROR parsing: %s", err)
		return
	}
	playload := params.Get("playload")
	fmt.Fprint(os.Stdout, playload)
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
