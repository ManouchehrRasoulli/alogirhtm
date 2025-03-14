package day7

import (
	"strconv"
)

type Record struct {
	FinalValue  int
	Numbers     []int
	IsReachable bool
	Operators   string
}

func MatchOperatorsForRecordPart1(records []Record) {
	for i, record := range records {
		ops := GenerateCombinationOfOperationInLen([]string{"+", "*"}, len(record.Numbers)-1, "")
		for _, op := range ops {
			v := DoOpsWithNumbers(op, record.Numbers)
			if record.FinalValue == v {
				records[i].IsReachable = true
				records[i].Operators = op
				break
			}
		}
	}
}

func MatchOperatorsForRecordPart2(records []Record) {
	for i, record := range records {
		ops := GenerateCombinationOfOperationInLen([]string{"+", "*", "|"}, len(record.Numbers)-1, "")
		for _, op := range ops {
			v := DoOpsWithNumbers(op, record.Numbers)
			if record.FinalValue == v {
				records[i].IsReachable = true
				records[i].Operators = op
				break
			}
		}
	}
}

func DoOpsWithNumbers(ops string, numbers []int) int {
	v := numbers[0]
	for i, o := range ops {
		switch o {
		case '+':
			v += numbers[i+1]
		case '*':
			v *= numbers[i+1]
		case '|':
			strA := strconv.Itoa(v)
			strB := strconv.Itoa(numbers[i+1])
			v, _ = strconv.Atoi(strA + strB)
		}
	}

	return v
}
