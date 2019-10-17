package piscine

import "fmt"

func PrintStr(str string) {
	charOneByOne := []byte(str)
	for _, letter := range charOneByOne {
		fmt.Println("%v", letter)
	}
}
