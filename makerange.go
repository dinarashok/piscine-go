package piscine

func MakeRange(min, max int) []int {
	var size int
	size = max - min
	var MakeSlice []int

	if min < max {
		MakeSlice = make([]int, size)
		for i := min; i < max; i++ {
			MakeSlice[i] = min + 1
			return MakeSlice
		}
	} else {
		var MakeSlice []int
		return MakeSlice
	}
	return MakeSlice
}
