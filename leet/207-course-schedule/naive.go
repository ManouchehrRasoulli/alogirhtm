package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) < 2 {
		return true
	}

	cMap := make(map[int]*course)
	for i := 0; i < len(prerequisites); i++ {
		var pr *course
		if pre, preOk := cMap[prerequisites[i][1]]; preOk {
			pr = pre
		} else {
			pr = &course{
				stat: notVisited,
				cid:  prerequisites[i][1],
				pre:  make([]*course, 0),
			}
			cMap[prerequisites[i][1]] = pr
		}

		var cr *course
		if cro, crOk := cMap[prerequisites[i][0]]; crOk {
			cr = cro
		} else {
			cr = &course{
				stat: notVisited,
				cid:  prerequisites[i][0],
				pre:  make([]*course, 0),
			}
			cMap[prerequisites[i][0]] = cr
		}

		cr.pre = append(cr.pre, pr)
	}

	for i := 0; i < numCourses; i++ {
		c := cMap[i]
		if hasCycle(c) {
			return false
		}
	}

	return true
}

const (
	notVisited = iota
	visiting
	visited
)

type course struct {
	stat int
	cid  int
	pre  []*course
}

func hasCycle(c *course) bool {
	if c == nil {
		return false
	}

	if c.stat == visiting {
		return true
	}

	if c.stat == visited {
		return false
	}

	c.stat = visiting

	for _, p := range c.pre {
		if hasCycle(p) {
			return true
		}
	}

	c.stat = visited

	return false
}
