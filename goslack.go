package main

import (
	"os"
	"fmt"
	"net"
	"flag"
	"time"

	"github.com/bsphere/le_go"
	"github.com/parnurzeal/gorequest"
	"github.com/tatsushid/go-fastping"
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
		fmt.Println("please provide the -slackpath parameter")
		os.Exit(1)
	}
	if letoken == "" {
		fmt.Println("please provide the -letoken parameter")
		os.Exit(1)
	}
	if text == "" {
		fmt.Println("please provide the -text parameter")
		os.Exit(1)
	}

	le, e := le_go.Connect(letoken)
	if e != nil {
		panic(e)
	}
	defer le.Close()
	le.Println(text)

	request := gorequest.New()
	_, _, err := request.Post(fmt.Sprintf("https://hooks.slack.com/services/%s", slackpath)).
	Set("User-Agent", "packethost/goslack").
	Send(`{"channel":"`+ channel +`", "username":"` + username + `", "text":"`+ text +`", "icon_emoji":":`+ emoji +`:"}`).
	End(printStatus)

	if err != nil {
		le.Printf("slack ping failed: %v", err)
	}

	p := fastping.NewPinger()

	v4, e := net.ResolveIPAddr("ip4:icmp", "147.75.192.73")
	if e == nil {
		p.AddIPAddr(v4)
	} else {
		le.Printf("ip4 147.75.192.73 error: %s", e)
	}

	v4priv, e := net.ResolveIPAddr("ip4:icmp", "10.100.0.73")
	if e == nil {
		p.AddIPAddr(v4priv)
	} else {
		le.Printf("ip4 10.100.0.73 ping error: %s", e)
	}

	v6, e := net.ResolveIPAddr("ip6:icmp", "2604:1380::49")
	if e == nil {
		p.AddIPAddr(v6)
	} else {
		le.Printf("ip6 2604:1380::49 error: %s", e)
	}

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		le.Printf("ping success: %s receive, RTT: %v\n", addr.String(), rtt)
	}

	e = p.Run()
	if e != nil {
		le.Printf("ping test error: %s", e)
	}
}

func printStatus(resp gorequest.Response, body string, errs []error){
	fmt.Println(resp.Status)
}
