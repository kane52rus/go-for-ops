package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	stringValue := "Старт"
	fmt.Println(stringStat(stringValue))
}

func stringStat(word string) string {
	resultMap := make(map[string]int)
	var keyArray []string
	var resultValue string
	for _, v := range word {
		letter := strings.ToLower(string(v))
		if letter == " " {
			continue
		} else if !slices.Contains(keyArray, letter) {
			keyArray = append(keyArray, letter)
			resultMap[letter]++
		} else {
			resultMap[letter]++
		}
	}
	for _, v := range keyArray {
		key := v
		value := strconv.Itoa(resultMap[v])
		if resultValue == "" {
			resultValue = resultValue + key + " - " + value
		} else {
			resultValue = resultValue + "\n"
			resultValue = resultValue + key + " - " + value
		}
	}
	return resultValue
}
