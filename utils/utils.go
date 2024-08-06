package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Specials = []string{"@", "#", "$", "%", "^", "&", "*"}
var Nums = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func IsPassCorrect(s string) bool {
	if len(s) >= 8 {
		numFlag := 0
		SpecialFlag := 0
		for _, i := range Nums {
			if strings.Contains(s, i) {
				numFlag++
			}
		}
		for _, i := range Specials {
			if strings.Contains(s, i) {
				SpecialFlag++
			}
		}
		if numFlag != 0 && SpecialFlag != 0 {
			return true
		} else {
			return false
		}
	}
	return false
}
