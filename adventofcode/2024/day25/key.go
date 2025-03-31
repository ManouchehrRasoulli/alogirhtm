package day25

import "fmt"

type Key struct {
	combination []int
}

func NewKey(combination []int) *Key {
	return &Key{combination: combination}
}

func (k *Key) String() string {
	return fmt.Sprintf("🗝️ %v", k.combination)
}
