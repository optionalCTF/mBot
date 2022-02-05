package main

import (
	"flag"
	"log"
	"time"

	"github.com/un4gi/mBot/config"
	"github.com/un4gi/mBot/env"
	"github.com/un4gi/mBot/mission"
	"github.com/un4gi/mBot/requests"
	"github.com/un4gi/mBot/targets"
)

func main() {
	token := flag.String("t", "", "Authorization: Bearer token")
	delay := flag.Uint("d", 1, "Time (in seconds) between requests")
	//edit := flag.Bool("m", false, "Edit mission with test")

	flag.Parse()

	requests.Token = *token
	config.Delay = *delay

	for {
		log.Printf(env.InfoColor, "Checking in...")
		targets.CheckTargets(requests.Urls[0])
		if config.LoggedIn {
			mission.CheckMissions(requests.Urls[1])
			targets.CheckForQR(requests.Urls[2])
		}

		secs := time.Duration(*delay) * time.Second
		time.Sleep(secs)
	}
}
