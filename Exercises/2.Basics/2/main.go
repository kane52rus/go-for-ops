package main

import (
	"fmt"
)

func main() {
	listOfString := []string{"a", "ab", "b", "cc", "bc", "c", "cd"}
	fmt.Println(isSorted(listOfString))
}

func isSorted(ww []string) bool {
	if len(ww) == 0 {
		return false
	}
	for k, v := range ww {
		if k == 0 {
			continue
		} else {
			currentValue := v
			previousValue := ww[k-1]
			currentValueLength := len(currentValue)
			previousValueLength := len(previousValue)
			var maxValue, minValue string
			if currentValueLength > previousValueLength {
				maxValue = currentValue
				minValue = previousValue
			} else {
				minValue = currentValue
				maxValue = previousValue
			}
			for m, _ := range minValue {
				minValueLetter := minValue[m]
				maxValueLetter := maxValue[m]
				if minValueLetter >= maxValueLetter {
					continue
				} else if currentValueLength > previousValueLength {
					continue
				} else {
					return false
				}
			}
		}
	}
	return true
}
