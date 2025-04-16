package day4

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	tickets := ReadTickets(data)

	total := 0
	for _, ticket := range tickets {
		total += ticket.PointsPart1()
	}

	t.Log(total) // 13
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	tickets := ReadTickets(data)

	total := 0
	for _, ticket := range tickets {
		total += ticket.PointsPart1()
	}

	t.Log(total) // 19135
}

func TestInputTestPart2(t *testing.T) {
	data, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	var (
		tickets    = ReadTickets(data)
		winCopies  = make(map[int]int)
		lastTicket = 0
		final      = 0
	)

	for _, ticket := range tickets {
		winCopies[ticket.TicketId]++
		if ticket.TicketId > lastTicket {
			lastTicket = ticket.TicketId
		}
	}

	for _, ticket := range tickets {
		ticket.PointsPart2(winCopies, lastTicket)
	}

	for _, ticket := range tickets {
		final += winCopies[ticket.TicketId]
	}

	t.Log(final) // 30
}

func TestInputPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	var (
		tickets    = ReadTickets(data)
		winCopies  = make(map[int]int)
		lastTicket = 0
		final      = 0
	)

	for _, ticket := range tickets {
		winCopies[ticket.TicketId]++
		if ticket.TicketId > lastTicket {
			lastTicket = ticket.TicketId
		}
	}

	for _, ticket := range tickets {
		ticket.PointsPart2(winCopies, lastTicket)
	}

	for _, ticket := range tickets {
		final += winCopies[ticket.TicketId]
	}

	t.Log(final) // 5704953
}
