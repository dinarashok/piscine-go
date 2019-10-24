package piscine

func AppendRange(min, max int) []int {
	var SliceResult []int
	if min < max {
		for i := min; i < max; i++ {

			SliceResult = append(SliceResult, i)
		}
	} else {
		return SliceResult
	}
	return SliceResult
}
