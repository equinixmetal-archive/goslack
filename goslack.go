package main

import (
	"os"
	"fmt"
	"flag"

	"github.com/parnurzeal/gorequest"
)

var channel, username, slackpath, text, emoji string
func init() {
	flag.StringVar(&slackpath, "slackpath", "",            "the path of the slack webhook")
	flag.StringVar(&text,      "text",      "",            "the message to post")
	flag.StringVar(&channel,   "channel",   "#general",    "the channel to post to")
	flag.StringVar(&username,  "username",  "goslackgo",   "the username")
	flag.StringVar(&emoji,     "emoji",     "poop",            "the empoji icon code without the colons")
}

func main() {
	flag.Parse()

	if slackpath == "" {
		fmt.Printf("please provide the -slackpath parameter\n")
		os.Exit(1)
	}
	if text == "" {
		fmt.Printf("please provide the -text parameter\n")
		os.Exit(1)
	}

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
