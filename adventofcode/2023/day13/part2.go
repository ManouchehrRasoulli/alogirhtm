package day13

func Part2(field Field) int {
	for horizontalMirror := 0; horizontalMirror < len(field)-1; horizontalMirror++ {
		var (
			diffCount = 0
		)

		for realDown := horizontalMirror + 1; realDown < len(field); realDown++ {
			shadowUp := (horizontalMirror - (realDown - horizontalMirror)) + 1
			down, downIsInMirror := field.StringRow(realDown)
			up, upIsInMirror := field.StringRow(shadowUp)
			if !downIsInMirror || !upIsInMirror {
				break
			}

			for i := range down {
				if down[i] != up[i] {
					diffCount++
				}
			}

			if diffCount > 1 {
				break
			}
		}

		if diffCount == 1 {
			return 100 * (horizontalMirror + 1)
		}
	}

	for verticalMirror := 0; verticalMirror < len(field[0])-1; verticalMirror++ {
		var (
			diffCount = 0
		)

		for realRight := verticalMirror + 1; realRight < len(field[0]); realRight++ {
			shadowLeft := (verticalMirror - (realRight - verticalMirror)) + 1
			right, rightIsInMirror := field.StringColumn(realRight)
			left, leftIsInMirror := field.StringColumn(shadowLeft)
			if !rightIsInMirror || !leftIsInMirror {
				break
			}

			for i := range right {
				if right[i] != left[i] {
					diffCount++
				}
			}

			if diffCount > 1 {
				break
			}
		}

		if diffCount == 1 {
			return verticalMirror + 1
		}
	}

	return 0
}
