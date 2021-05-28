package main

import "fmt"

func main() {
	str := "Liam"
	reversedStr := reverseString(str)
	fmt.Println(reversedStr)
}

func reverseString(str string) string {
	var reversedStr string
	for _, s := range str {
		reversedStr = string(s) + reversedStr
	}
	return reversedStr
}
