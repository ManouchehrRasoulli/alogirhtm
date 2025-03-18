package day14

import (
	"log"
	"strings"
)

func ParseBots(data string) []*Bot {
	var (
		bots = make([]*Bot, 0)
	)

	lines := strings.Split(data, "\n")
	for num, line := range lines {
		bot, err := NewBotFromString(line)
		if err != nil {
			log.Printf("Invalid data line %d: %s, %v\n", num, line, err)
			continue
		}

		bots = append(bots, bot)
	}

	return bots
}
