package day5

import (
	"strconv"
	"strings"
)

type Rule struct {
	Precedence int
	Item       int
}

type Update struct {
	Items   []int
	IsValid bool

	// if invalid because of rule
	Cause []string
}

func Parse(input string) ([]Rule, []Update) {
	var (
		rules   []Rule
		updates []Update
	)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.Contains(line, "|") { // rule
			seq := strings.Split(line, "|")
			precedence, _ := strconv.Atoi(seq[0])
			item, _ := strconv.Atoi(seq[1])
			rules = append(rules, Rule{
				Precedence: precedence,
				Item:       item,
			})
		} else { // updates
			update := Update{
				Items:   nil,
				IsValid: true,
				Cause:   make([]string, 0),
			}

			items := strings.Split(line, ",")
			for _, item := range items {
				itemInt, _ := strconv.Atoi(item)
				update.Items = append(update.Items, itemInt)
			}

			updates = append(updates, update)
		}
	}

	return rules, updates
}

type PrecedenceHierarchy struct {
	mp map[int]map[int]struct{}
}

func (p *PrecedenceHierarchy) Precede(item int, precede int) bool {
	if v, ok := p.mp[item]; !ok {
		return false
	} else {
		if _, ok = v[precede]; ok {
			return true
		}
	}

	return false
}

func CreatePrecedenceMap(rules []Rule) PrecedenceHierarchy {
	mp := make(map[int]map[int]struct{})
	for _, rule := range rules {
		if _, ok := mp[rule.Item]; !ok {
			mp[rule.Item] = make(map[int]struct{})
		}
		mp[rule.Item][rule.Precedence] = struct{}{}
	}

	return PrecedenceHierarchy{
		mp: mp,
	}
}
