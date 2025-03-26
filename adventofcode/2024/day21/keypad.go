package day21

import (
	"fmt"
	"log"
	"math"
)

var (
	keyPad = map[rune]map[rune][]rune{
		'A': {'A': {}, '^': {'<'}, '>': {'v'}, 'v': {'<', 'v'}, '<': {'v', '<', '<'}},
		'^': {'^': {}, 'A': {'>'}, '>': {'v', '>'}, 'v': {'v'}, '<': {'v', '<'}},
		'<': {'<': {}, 'v': {'>'}, '^': {'>', '^'}, '>': {'>', '>'}, 'A': {'>', '>', '^'}},
		'v': {'v': {}, '<': {'<'}, '>': {'>'}, '^': {'^'}, 'A': {'^', '>'}},
		'>': {'>': {}, 'A': {'^'}, '^': {'<', '^'}, 'v': {'<'}, '<': {'<', '<'}}}
)

type Option func(k *Keypad)

func WithKeyPad(sequence string) Option {
	return func(k *Keypad) {
		keypad := NewKeypad(WithNumPad(sequence))
		k.keyPad = keypad
	}
}

func WithNumPad(sequence string) Option {
	return func(k *Keypad) {
		numpad := NewNumpad(sequence)
		k.numPad = numpad
	}
}

type Keypad struct {
	numPad *Numpad // have other numpad or other keypad
	keyPad *Keypad
}

func NewKeypad(options ...Option) *Keypad {
	k := Keypad{}

	for _, opt := range options {
		opt(&k)
	}

	if k.keyPad == nil && k.numPad == nil {
		log.Fatal("one option required !!")
	}

	if k.keyPad != nil && k.numPad != nil {
		log.Fatal("only one option is valid !!")
	}

	return &k
}

func (k *Keypad) KeyToDirections() (string, [][]rune) {
	var (
		subPaths      = make([][]rune, 0)
		paths         = make([][]rune, 0)
		computedPaths = make(map[string]struct{})
		seq           string
		minLen        = math.MaxInt
	)

	if k.keyPad != nil {
		seq, subPaths = k.keyPad.KeyToDirections()
	} else if k.numPad != nil {
		seq, subPaths = k.numPad.NumsToDirections()
	}

	for _, subPath := range subPaths {
		ip := make([]rune, 0)
		ip = append(ip, confirm)
		ip = append(ip, subPath...)
		paths = append(paths, ip)
	}

	for _, path := range paths {
		internalPaths := make([][]rune, 0)
		for i := 1; i < len(path); i++ {
			ac, bc := path[i-1], path[i]
			directions := keypadPathFromAtoB(ac, bc)
			internalPaths = cartesianProduct(internalPaths, directions)
		}

		for _, internalPath := range internalPaths {
			computedPaths[string(internalPath)] = struct{}{}
		}
	}

	final := make([][]rune, 0)
	for key, _ := range computedPaths {
		if len(key) < minLen {
			final = make([][]rune, 0)
			minLen = len(key)
			final = append(final, []rune(key))
		} else if len(key) == minLen {
			final = append(final, []rune(key))
		}
	}

	return seq, final
}

func keypadPathFromAtoB(a rune, b rune) [][]rune {
	var (
		key    = fmt.Sprintf("%s-%s", string(a), string(b))
		runes  = keyPad[a][b]
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
