package time_bucket

import (
	"time"
)

const (
	maxCount = 11
)

type bucket struct {
	responseTime time.Duration
	count        int
	next         *bucket
}

func appendTo(head *bucket, count int, responseTime time.Duration) *bucket {
	if count > maxCount {
		// already reached the maximum list len
		return nil
	}

	if head == nil {
		return &bucket{
			responseTime: responseTime,
			count:        1,
			next:         nil,
		}
	}

	if head.responseTime == responseTime {
		head.count++
		return head
	}

	if head.responseTime < responseTime {
		next := appendTo(head.next, count+1, responseTime)
		head.next = next
		return head
	}

	return &bucket{
		responseTime: responseTime,
		count:        1,
		next:         head,
	}
}

var (
	head = &bucket{
		responseTime: 0,
		count:        0,
	}
)

func CaptureResponseTime(responseTime time.Duration) {
	head = appendTo(head, 0, responseTime)
}

func ListResponseTimes() []time.Duration {
	var (
		cursor = head.next
		list   []time.Duration
	)

	for i := 1; i < maxCount; i++ {
		if cursor != nil {
			list = append(list, cursor.responseTime)
			cursor = cursor.next
			continue
		}

		break
	}

	if cursor != nil {
		// remove items after first max params
		cursor.next = nil
	}

	return list
}
