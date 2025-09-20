package change

func greedyChangeMoney(amount int, minAmount int, maxAmount int) (buckets []int) {
	if amount < minAmount {
		return buckets
	}

	maxBucketCount := amount / maxAmount
	reminder := amount % maxAmount

	for i := 0; i < maxBucketCount; i++ {
		buckets = append(buckets, maxAmount)
	}

	if reminder == 0 {
		return buckets
	}

	if reminder > minAmount {
		buckets = append(buckets, reminder)
		return buckets
	}

	diff := minAmount - reminder
	buckets[len(buckets)-1] -= diff
	reminder += diff
	buckets = append(buckets, reminder)
	return buckets
}
