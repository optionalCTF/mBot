package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/un4gi/mBot/config"
	"github.com/un4gi/mBot/discord"
	"github.com/un4gi/mBot/env"
	"github.com/un4gi/mBot/mission"
	"github.com/un4gi/mBot/requests"
	"github.com/un4gi/mBot/targets"
)

func main() {
	token := flag.String("t", "", "Authorization: Bearer token")
	delay := flag.Uint("d", 60, "Time (in seconds) between requests")
	disc := flag.Bool("n", false, "Connect to Discord?")

	flag.Parse()

	// Changed expected result for automated run
	if *token != "config" {
		fmt.Printf(env.DebugColor, "You need to supply an Authorization: Bearer token.")
		os.Exit(0)
	} else {
		requests.Token = *token
	}

	if *disc {
		discord.ConnectDiscord("[~] Mission Bot Connection Established")
	}

	config.Delay = *delay
	targets.CheckTargets(requests.Urls[0])
	mission.CheckClaimed()
	for {
		log.Printf(env.InfoColor, "Checking in...")
		targets.CheckTargets(requests.Urls[0])
		if config.LoggedIn {
			targets.CheckForQR(requests.Urls[2])
			if mission.CheckWallet(requests.Urls[6]) {
				mission.CheckMissions(requests.Urls[1])
			}
		}

		secs := time.Duration(*delay) * time.Second
		time.Sleep(secs)
	}
}
