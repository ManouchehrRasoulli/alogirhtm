package day4

import (
	"strconv"
	"strings"
)

type Ticket struct {
	TicketId       int
	WinningNumbers []int
	YourNumbers    []int
}

func (t Ticket) PointsPart1() int {
	var (
		point int
	)

	for _, w := range t.WinningNumbers {
		for _, h := range t.YourNumbers {
			if w == h {
				if point == 0 {
					point = 1
				} else {
					point *= 2
				}
			}
		}
	}

	return point
}

func (t Ticket) MatchingNumbers() int {
	var (
		matches int
	)

	for _, w := range t.WinningNumbers {
		for _, h := range t.YourNumbers {
			if w == h {
				matches++
			}
		}
	}

	return matches
}

func (t Ticket) PointsPart2(winCopy map[int]int, lastTicket int) {
	ticketScore := t.MatchingNumbers()
	count := winCopy[t.TicketId]

	for i := 0; i < count; i++ {
		for j := t.TicketId + 1; j <= t.TicketId+ticketScore && j <= lastTicket; j++ { // win one copy of each cards
			winCopy[j]++
		}
	}
}

func ReadTickets(data string) []Ticket {
	var (
		lines   = strings.Split(data, "\n")
		tickets = make([]Ticket, 0)
	)

	for cardId, line := range lines {
		card := strings.Split(line, ":")
		nums := strings.Split(card[1], "|")

		tickets = append(tickets, Ticket{
			TicketId:       cardId + 1,
			WinningNumbers: ArrayStrToArrayInt(strings.Split(nums[0], " ")),
			YourNumbers:    ArrayStrToArrayInt(strings.Split(nums[1], " ")),
		})
	}

	return tickets
}

func ArrayStrToArrayInt(nums []string) []int {
	var (
		output = make([]int, 0)
	)

	for _, n := range nums {
		if n == "" {
			continue
		}

		num, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		output = append(output, num)
	}

	return output
}
