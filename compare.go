package piscine

func Compare(a, b string) int {
	RuneA := []rune(a)
	RuneB := []rune(b)

	// find lenths of those massives
	lenA := 0
	lenB := 0
	for range RuneA {
		lenA++
	}
	for range RuneB {
		lenB++
	}

	for i := 0; i < lenA-lenB; i++ {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
	}
	return -1
}
