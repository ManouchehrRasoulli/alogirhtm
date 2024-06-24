package _kyu

import "testing"

func TestItem(t *testing.T) {
	head := item{
		value: 10,
		next: &item{
			value: 11,
			next: &item{
				value: 9,
				next:  nil,
			},
		},
	}

	t.Log(head.String())

	shiftPlace(&head)

	t.Log(head.String())
}

func TestSolution(t *testing.T) {
	t.Log(Solution0([]int{1, 21, 55}))
	t.Log(Solution0([]int{6, 9, 21}))
	t.Log(Solution0([]int{9}))
	t.Log(Solution0([]int{1, 1, 2, 2, 2, 1, 1}))
}

func TestDivides(t *testing.T) {
	t.Log(9 % 6)
	t.Log(21 % 9)
	t.Log(3 % 2)
	t.Log(2 % 1)
	t.Log(13 % 9)
}
