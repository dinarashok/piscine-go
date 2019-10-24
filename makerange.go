package piscine

func MakeRange(min, max int) []int {
	var MakeSlice []int
	if max > min {
		MakeSlice = make([]int, max-min)
		for i := 0; i < max-min; i++ {
			MakeSlice[i] = min + i
		}
	} else {
		var MakeSlice []int
		return MakeSlice
	}
	return MakeSlice
}
