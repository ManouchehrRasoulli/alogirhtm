package day21

import (
	"fmt"
)

const (
	noValue = ' '
	confirm = 'A'
)

var (
	cache = make(map[string][][]rune)
)

var (
	numPad = map[rune]map[rune][]rune{
		'7': {'7': {}, '8': {'>'}, '9': {'>', '>'},
			'4': {'v'}, '5': {'v', '>'}, '6': {'v', '>', '>'},
			'1': {'v', 'v'}, '2': {'v', 'v', '>'}, '3': {'v', 'v', '>', '>'},
			'0': {'>', 'v', 'v', 'v'}, 'A': {'>', '>', 'v', 'v', 'v'}},

		'8': {'7': {'<'}, '8': {}, '9': {'>'},
			'4': {'<', 'v'}, '5': {'v'}, '6': {'v', '>'},
			'1': {'<', 'v', 'v'}, '2': {'v', 'v'}, '3': {'v', 'v', '>'},
			'0': {'v', 'v', 'v'}, 'A': {'v', 'v', 'v', '>'}},

		'9': {'7': {'<', '<'}, '8': {'<'}, '9': {},
			'4': {'<', '<', 'v'}, '5': {'<', 'v'}, '6': {'v'},
			'1': {'<', '<', 'v', 'v'}, '2': {'<', 'v', 'v'}, '3': {'v', 'v'},
			'0': {'<', 'v', 'v', 'v'}, 'A': {'v', 'v', 'v'}},

		'4': {'7': {'^'}, '8': {'^', '>'}, '9': {'^', '>', '>'},
			'4': {}, '5': {'>'}, '6': {'>', '>'},
			'1': {'v'}, '2': {'v', '>'}, '3': {'v', '>', '>'},
			'0': {'>', 'v', 'v'}, 'A': {'>', '>', 'v', 'v'}},

		'5': {'7': {'<', '^'}, '8': {'^'}, '9': {'^', '>'},
			'4': {'<'}, '5': {}, '6': {'>'},
			'1': {'<', 'v'}, '2': {'v'}, '3': {'v', '>'},
			'0': {'v', 'v'}, 'A': {'v', 'v', '>'}},

		'6': {'7': {'<', '<', '^'}, '8': {'<', '^'}, '9': {'^'},
			'4': {'<', '<'}, '5': {'<'}, '6': {},
			'1': {'<', '<', 'v'}, '2': {'<', 'v'}, '3': {'v'},
			'0': {'<', 'v', 'v'}, 'A': {'v', 'v'}},

		'1': {'7': {'^', '^'}, '8': {'^', '^', '>'}, '9': {'^', '^', '>', '>'},
			'4': {'^'}, '5': {'^', '>'}, '6': {'^', '>', '>'},
			'1': {}, '2': {'>'}, '3': {'>', '>'},
			'0': {'>', 'v'}, 'A': {'>', '>', 'v'}},

		'2': {'7': {'<', '^', '^'}, '8': {'^', '^'}, '9': {'>', '^', '^'},
			'4': {'<', '^'}, '5': {'^'}, '6': {'^', '>'},
			'1': {'<'}, '2': {}, '3': {'>'},
			'0': {'v'}, 'A': {'v', '>'}},

		'3': {'7': {'<', '<', '^', '^'}, '8': {'<', '^', '^'}, '9': {'^', '^'},
			'4': {'^', '<', '<'}, '5': {'<', '^'}, '6': {'^'},
			'1': {'<', '<'}, '2': {'<'}, '3': {},
			'0': {'<', 'v'}, 'A': {'v'}},

		'0': {'7': {'^', '^', '^', '<'}, '8': {'^', '^', '^'}, '9': {'^', '^', '^', '>'},
			'4': {'^', '^', '<'}, '5': {'^', '^'}, '6': {'^', '^', '>'},
			'1': {'^', '<'}, '2': {'^'}, '3': {'^', '>'},
			'0': {}, 'A': {'>'}},

		'A': {'7': {'^', '^', '^', '<', '<'}, '8': {'<', '^', '^', '^'}, '9': {'^', '^', '^'},
			'4': {'^', '^', '<', '<'}, '5': {'^', '^', '<'}, '6': {'^', '^'},
			'1': {'^', '<', '<'}, '2': {'^', '<'}, '3': {'^'},
			'0': {'<'}, 'A': {}}}
)

type Numpad struct {
	nums string
}

func NewNumpad(nums string) *Numpad {
	return &Numpad{
		nums: nums,
	}
}

func (n *Numpad) NumsToDirections() (string, [][]rune) {
	var (
		paths = make([][]rune, 0)
	)

	nums := string(confirm) + n.nums
	for i := 1; i < len(nums); i++ {
		ac, bc := nums[i-1], nums[i]
		directions := numpadPathsFromAtoB(rune(ac), rune(bc))
		paths = cartesianProduct(paths, directions)
	}

	return n.nums, paths
}

func cartesianProduct(first [][]rune, second [][]rune) [][]rune {
	cartesian := make([][]rune, 0)

	if len(first) == 0 {
		for _, b := range second {
			items := make([]rune, 0)
			items = append(items, b...)
			items = append(items, confirm)
			cartesian = append(cartesian, items)
		}

		return cartesian
	}

	if len(second) == 0 {
		for _, a := range first {
			items := make([]rune, 0)
			items = append(items, a...)
			items = append(items, confirm)
			cartesian = append(cartesian, items)
		}

		return cartesian
	}

	for _, a := range first {
		for _, b := range second {
			items := make([]rune, 0)
			items = append(items, a...)
			items = append(items, b...)
			items = append(items, confirm)
			cartesian = append(cartesian, items)
		}
	}

	return cartesian
}

func numpadPathsFromAtoB(a rune, b rune) [][]rune {
	var (
		key    = fmt.Sprintf("%s-%s", string(a), string(b))
		runes  = numPad[a][b]
		perms  = make([][]rune, 0)
		result = make(map[string]struct{})
	)

	if v, ok := cache[key]; ok {
		return v
	}

	permute(runes, 0, len(runes)-1, result)
	for k, _ := range result {
		perms = append(perms, []rune(k))
	}

	cache[key] = perms

	return perms
}

func permute(runes []rune, left int, right int, result map[string]struct{}) {
	if left == right {
		result[string(runes)] = struct{}{}
	} else {
		for i := left; i <= right; i++ {
			runes[left], runes[i] = runes[i], runes[left]
			permute(runes, left+1, right, result)
			runes[left], runes[i] = runes[i], runes[left]
		}
	}
}
