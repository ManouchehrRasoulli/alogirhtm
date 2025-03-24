package day19

func Match(input string, patterns []string) bool {
	var (
		n               = len(input)
		checkedSequence = make([]bool, n+1)
	)
	checkedSequence[0] = true // the first case means that the zero len sequence is matched

	for i := 1; i < n+1; i++ {
		checkedSequence[i] = false
		for _, pattern := range patterns {
			patternLen := len(pattern)
			if i >= patternLen && input[i-patternLen:i] == pattern {
				checkedSequence[i] = checkedSequence[i] || checkedSequence[i-patternLen]
			}
		}
	}

	return checkedSequence[n]
}
