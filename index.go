package piscine

func Index(s string, toFind string) int {
	RuneS := []rune(s)
	RuneToFind := []rune(toFind)
	lenS := 0
	lenToFind := 0
	for range RuneS {
		lenS++
	}
	for range RuneToFind {
		lenToFind++
	}
	for i := 0; i <= lenS-lenToFind; i++ {
		if toFind == s[i:i+lenToFInd] {
			return (i)
		}
	}
	return -1
}
