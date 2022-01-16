package main

import "fmt"

func main() {
	inputString := "kl☺fsdjpvo"
	charToFind := 'd'
	fmt.Printf("First position of char %c is %d",
		charToFind, findPos2(inputString, charToFind))
}
func findPosWrong(s string, r rune) int {
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == r {
			return i
		}
	}
	return -1
}
func findPos(s string, r rune) int {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] == r {
			return i
		}
	}
	return -1
}
func findPos2(s string, r rune) int {
	for i, c := range []rune(s) {
		if c == r {
			return i
		}
	}
	return -1
}
