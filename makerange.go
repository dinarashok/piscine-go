package piscine

func MakeRange(min, max int) []int {
	var size int
	size = max - min

	if min < max {
		MakeSlice := make([]int, size)
		for i := min; i < max; i++ {
			MakeSlice[i] = min + 1
		}
	} else {
		var MakeSlice []int
		return MakeSlice
	}
	return MakeSlice
}
