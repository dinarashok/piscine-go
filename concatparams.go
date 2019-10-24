package piscine

func ConcatParams(args []string) string {
	len := 0
	for range args {
		len++
	}

	var answer string
	for i := range args {
		if i != lenargs-1 {
			answer = answer + args[i] + "\n"
		} else {
			answer = answer + args[i]
		}
	}
	return answer
}
