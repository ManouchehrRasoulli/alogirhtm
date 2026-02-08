package day11

import (
	_ "embed"
)

//go:embed sample1.txt
var sample1 string

//go:embed sample2.txt
var sample2 string

//go:embed puzzle.txt
var puzzle string
