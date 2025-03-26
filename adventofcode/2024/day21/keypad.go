package day21

import (
	"log"
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

func (k *Keypad) KeyToDirections() []rune {
	var (
		path         = make([]rune, 0)
		internalPath = make([]rune, 1)
	)

	internalPath[0] = confirm

	if k.keyPad != nil {
		internalPath = append(internalPath, k.keyPad.KeyToDirections()...)
	} else if k.numPad != nil {
		internalPath = append(internalPath, k.numPad.NumsToDirections()...)
	}

	log.Println("keypad compute directions ::", string(internalPath))

	for i := 1; i < len(internalPath); i++ {
		ac, bc := internalPath[i-1], internalPath[i]
		path = append(path, keypadPathFromAtoB(ac, bc)...)
		path = append(path, confirm)
	}

	log.Println("keypad output ::", string(path))

	return path
}

func keypadPathFromAtoB(a rune, b rune) []rune {
	return keyPad[a][b]
}
