package _kyu

import (
	"fmt"
	"slices"
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

	arc := make([]int, len(ar), len(ar))
	for i, v := range ar {
		arc[i] = v
	}

	slices.SortFunc(arc, func(a, b int) int {
		if a < b {
			return +1
		}

		return -1
	})

	return smallestDifInRange(arc, 0, len(arc)-1) * len(arc)
}

func min_(l, r int) (min, max int) {
	if l < r {
		return l, r
	}

	return r, l
}

func di_(mx, mi int) int {
	if mi == mx {
		return mi
	}

	di := mx % mi
	if di == 0 {
		return mi
	}

	return di
}

func smallestDifInRange(ar []int, i, j int) int {
	if i == j {
		return ar[i]
	}

	if j-i != 1 {
		d := (j - i) / 2
		l := smallestDifInRange(ar, i, i+d)
		r := smallestDifInRange(ar, i+d+1, j)
		mi, mx := min_(l, r)
		return di_(mx, mi)
	}

	mi, mx := min_(ar[i], ar[j])
	return di_(mx, mi)
}
