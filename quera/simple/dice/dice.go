package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var x string
	_, err := fmt.Scanln(&x)
	if err != nil {
		log.Println("got error reading input.txt", err)
		os.Exit(-1)
	}

	var num int
	num, err = strconv.Atoi(x)
	if err != nil {
		log.Println("non integer input.txt", err, x)
		os.Exit(-1)
	}

	switch num {
	case 1:
		fmt.Println(6)
	case 2:
		fmt.Println(5)
	case 3:
		fmt.Println(4)
	case 4:
		fmt.Println(3)
	case 5:
		fmt.Println(2)
	default:
		fmt.Println(1)
	}
}
