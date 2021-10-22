package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"mBot/env"
	"mBot/mission"
	"mBot/targets"
)

func main() {
	token := flag.String("t", "", "Authorization: Bearer token")
	delay := flag.Uint("d", 60, "Time (in seconds) between requests")

	flag.Parse()

	urls := []string{
		// urls[0] = all unregistered targets
		"https://platform.synack.com/api/targets?filter%5Bprimary%5D=unregistered&filter%5Bsecondary%5D=all&filter%5Bcategory%5D=all&filter%5Bindustry%5D=all&sorting%5Bfield%5D=dateUpdated&sorting%5Bdirection%5D=desc",
		// urls[1] = available missions sorted by price
		"https://platform.synack.com/api/tasks/v1/tasks?sortBy=price-sort-desc&withHasBeenViewedInfo=true&status=PUBLISHED&page=0&pageSize=20",
		// urls[2] = QR window
		"https://platform.synack.com/api/targets?filter%5Bprimary%5D=all&filter%5Bsecondary%5D%5B%5D=a&filter%5Bsecondary%5D%5B%5D=l&filter%5Bsecondary%5D%5B%5D=l&filter%5Bsecondary%5D%5B%5D=quality_period&filter%5Bcategory%5D=all&filter%5Bindustry%5D=all&sorting%5Bfield%5D=dateUpdated&sorting%5Bdirection%5D=desc",
		// urls[3] = claimed missions
		"https://platform.synack.com/api/tasks/v1/tasks?withHasBeenViewedInfo=true&status=CLAIMED&page=0&pageSize=20",
		// urls[4] = beginning of URL to edit missions
		"https://platform.synack.com/api/tasks/v1/organizations/",
	}

	headers := []string{"mBot/1.0", "Bearer " + *token, "same-origin", "cors", "https://platform.synack.com/tasks/user/available", "xxxx", "application/json", "close"}

	if len(*token) == 0 {
		fmt.Printf(env.DebugColor, "You need to supply an Authorization: Bearer token.")
		os.Exit(0)
	}

	for {
		log.Printf(env.InfoColor, "Checking in...")
		targets.CheckTargets(urls[0], headers)
		mission.CheckMissions(urls[1], headers)
		targets.CheckForQR(urls[2], headers)

		secs := time.Duration(*delay) * time.Second
		time.Sleep(secs)
	}

}
