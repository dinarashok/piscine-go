package printcomb2

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			b := j + 1
			for a := i; a <= '9'; a++ {
				for ; b <= '9'; b++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(a)
					z01.PrintRune(b)
					if i != '9' || j != '8' || a != '9' || b != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				b = '0'
			}
		}
	}
	z01.PrintRune(10)
}

func main() {
	PrintComb2()
}
