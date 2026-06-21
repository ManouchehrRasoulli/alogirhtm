package day13

func Part1(field Field) int {
	// horizontal
	for horizontalMirror := 0; horizontalMirror < len(field)-1; horizontalMirror++ {
		mirrorFound := true

		for realDown := horizontalMirror + 1; realDown < len(field); realDown++ {
			shadowUp := (horizontalMirror - (realDown - horizontalMirror)) + 1
			down, downIsInMirror := field.StringRow(realDown)
			up, upIsInMirror := field.StringRow(shadowUp)
			if !downIsInMirror || !upIsInMirror {
				break
			}
			if down != up {
				mirrorFound = false
				break
			}
		}

		if mirrorFound {
			return 100 * (horizontalMirror + 1)
		}
	}

	// vertical
	for verticalMirror := 0; verticalMirror < len(field[0])-1; verticalMirror++ {
		mirrorFound := true

		for realRight := verticalMirror + 1; realRight < len(field[0]); realRight++ {
			shadowLeft := (verticalMirror - (realRight - verticalMirror)) + 1
			right, rightIsInMirror := field.StringColumn(realRight)
			left, leftIsInMirror := field.StringColumn(shadowLeft)
			if !rightIsInMirror || !leftIsInMirror {
				break
			}
			if right != left {
				mirrorFound = false
				break
			}
		}

		if mirrorFound {
			return verticalMirror + 1
		}
	}

	return 0
}
