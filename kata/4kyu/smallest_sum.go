package _kyu

import (
	"fmt"
)

type item struct {
	value int
	next  *item
}

func (i *item) String() string {
	if i == nil {
		return "end"
	}
	return fmt.Sprintf("%d -> %s", i.value, i.next.String())
}

func (i *item) areAllSame() bool {
	if i.next == nil {
		return true
	}

	if i.value != i.next.value {
		return false
	}

	return i.next.areAllSame()
}

// shift item to right place
func shiftPlace(i *item) {
	if i.next == nil {
		return
	}

	if i.next.value > i.value {
		nvc := i.next.value
		i.next.value = i.value
		i.value = nvc
	}

	shiftPlace(i.next)
}

func Solution0(ar []int) int {
	if len(ar) == 1 {
		return ar[0]
	}

	var head *item

	for _, v := range ar {
		n := item{
			value: v,
			next:  head,
		}

		if head == nil {
			head = &n
			continue
		}

		n.next = head
		shiftPlace(&n)
		head = &n
	}

	shiftPlace(head)

	for !head.areAllSame() {
		// time.Sleep(time.Second)
		p := head
		for p.next != nil && p.value == p.next.value {
			p = p.next
		}
		p.value = p.value - p.next.value
		shiftPlace(head)
		// fmt.Println(head.String())
	}

	// fmt.Println(head.areAllSame())

	return head.value * len(ar)
}

func Solution(ar []int) int {
	if len(ar) == 1 {
		return ar[0]
	}

	return smallestInRange(ar, 0, len(ar)-1) * len(ar)
}

func smallestInRange(ar []int, i, j int) int {
	if i == j {
		return ar[i]
	}

	if j-i != 1 {
		d := (j - i) / 2
	}
}
