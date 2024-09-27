package _kyu

import (
	"fmt"
	"strings"
)

func Capitalize(st string, arr []int) string {
	fmt.Println(st)
	stN := strings.Builder{}
	arrIndex := 0
	for i, v := range st {
		if arrIndex < len(arr) && arr[arrIndex] == i {
			stN.WriteByte(byte(v - 32))
			arrIndex++
			continue
		}
		stN.WriteByte(byte(v))
	}

	return stN.String()
}
