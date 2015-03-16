package main

import (
	"os"
	"fmt"
	"flag"

	"github.com/bsphere/le_go"
	"github.com/parnurzeal/gorequest"
)

var slackpath, letoken, text, channel, username, emoji string
func init() {
	flag.StringVar(&slackpath, "slackpath", "",          "the path of the slack webhook")
	flag.StringVar(&letoken,   "letoken",   "",          "the log entry token")
	flag.StringVar(&text,      "text",      "",          "the message to post")
	flag.StringVar(&channel,   "channel",   "#general",  "the channel to post to")
	flag.StringVar(&username,  "username",  "goslackgo", "the username")
	flag.StringVar(&emoji,     "emoji",     "poop",      "the empoji icon code without the colons")
}

func main() {
	flag.Parse()

	if slackpath == "" {
		fmt.Printf("please provide the -slackpath parameter\n")
		os.Exit(1)
	}
	if letoken == "" {
		fmt.Printf("please provide the -letoken parameter\n")
		os.Exit(1)
	}
	if text == "" {
		fmt.Printf("please provide the -text parameter\n")
		os.Exit(1)
	}

	le, e := le_go.Connect(letoken)
	if e != nil {
		panic(e)
	}
	defer le.Close()
	le.Println(text)

	request := gorequest.New()
	_, body, err := request.Post(fmt.Sprintf("https://hooks.slack.com/services/%s", slackpath)).
	Set("User-Agent", "packethost/goslack").
	Send(`{"channel":"`+ channel +`", "username":"` + username + `", "text":"`+ text +`", "icon_emoji":":`+ emoji +`:"}`).
	End(printStatus)

	if err != nil {
		fmt.Sprintf("Error: %v", err)
	}	else {
		fmt.Println(body)
	}
}

func printStatus(resp gorequest.Response, body string, errs []error){
	fmt.Println(resp.Status)
}
