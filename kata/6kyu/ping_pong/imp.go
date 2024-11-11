package kata

import (
	"strings"
)

func isNoise(sound string) bool {
	return sound != "ping" && sound != "pong"
}

func addScore(player string, aScore, bScore *int) {
	switch player {
	case "ping":
		*aScore++
		return
	case "pong":
		*bScore++
		return
	}
}

func PingPong(sounds string) string {
	acts := strings.Split(sounds, "-")

	afterSound := false

	roundStarter := acts[0]
	lastToucher := ""

	pingScore := 0
	pongScore := 0

	for i := 0; i < len(acts); i++ {
		if isNoise(acts[i]) {
			if !afterSound && roundStarter != lastToucher {
				addScore(roundStarter, &pingScore, &pongScore)
			}
			afterSound = true
			continue
		}

		lastToucher = acts[i]

		if afterSound {
			roundStarter = acts[i]
			afterSound = false
		}
	}

	if pingScore == pongScore {
		addScore(lastToucher, &pongScore, &pingScore)
	}

	if pingScore > pongScore {
		return "ping"
	}

	return "pong"
}
