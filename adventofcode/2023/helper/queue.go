package helper

type PriorityQueue struct {
	Queue []PriorityQueueElement
}

func (q *PriorityQueue) Push(newEl PriorityQueueElement) {
	for i, el := range q.Queue {
		if newEl.Value() <= el.Value() {
			// Insert the new element in the correct position
			q.Queue = append(q.Queue[:i], append([]PriorityQueueElement{newEl}, q.Queue[i:]...)...)
			return
		}
	}
	// If we get here, the new element has the lowest priority, so add it to the end
	q.Queue = append(q.Queue, newEl)
}

func (q *PriorityQueue) Len() int {
	return len(q.Queue)
}

func (q *PriorityQueue) Pop() *PriorityQueueElement {
	if q.Len() == 0 {
		return nil
	}
	el := q.Queue[0]
	q.Queue = q.Queue[1:]
	return &el
}

func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{[]PriorityQueueElement{}}
}

type PriorityQueueElement interface {
	Value() int
}

type PositionPriorityQueueElement struct {
	value    int
	Position Location
}

func NewPositionPriorityQueueElement(value int, position Location) PositionPriorityQueueElement {
	return PositionPriorityQueueElement{value, position}
}

func (el PositionPriorityQueueElement) Value() int {
	return el.value
}
