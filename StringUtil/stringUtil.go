package stringUtil

import (
	"fmt"
	"time"
)

func CommonString(a, b string) string {
	start := time.Now()

	defer fmt.Println(time.Now().Format(time.RFC850))

	solution := ""
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[0] {
			for j := 0; j < len(b); j++ {

				if a[i+j] != b[j] {
					count = 0
					break
				} else {
					count++
					solution += string(a[i+j])
				}
			}
			if count == len(b) {
				break
			} else {
				solution = ""
			}
		}
	}
	end := time.Since(start)
	fmt.Println(end)

	return solution
}
