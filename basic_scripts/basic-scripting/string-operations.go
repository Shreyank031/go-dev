package main

import (
	"fmt"
	"strings"
)

func main() {
	//Concatinate in go
	str1 := "Hello"
	str2 := "Go"
	concatinate := str1 + " " + str2
	fmt.Println("Concatinated strings -> ", concatinate)

	//split
	sentence := "Go is an Open Source Language"
	split := strings.Split(sentence, "")
	splitt := strings.Split(sentence, " ")
	fmt.Println("\n", split, "\n", splitt)

	//Replace a substring
	replaceSubstring := strings.Replace(sentence, "Go", "Golang", 1)
	fmt.Println("\n", replaceSubstring)

	// Uppercase and Lowercase
	upper := strings.ToLower(sentence)
	lower := strings.ToUpper(sentence)
	fmt.Println("\n", upper, "\n", lower)

	//Trim spaces
	whiteSpaces := "       trim white spaces   "
	trimSpaces := strings.TrimSpace(whiteSpaces)
	fmt.Println("\n", trimSpaces)

	//Substring using slicing

	start := 3
	end := 13

	subString := sentence[start:end]
	subStringg := sentence[8:20]
	fmt.Println("\n", subString)
	fmt.Println(subStringg)

	// check if the string contains a substring
	contains := strings.Contains(sentence, "Source") //output is boolean -> true
	fmt.Println()
	fmt.Println(contains)

	// Finding the substirng index
	index := strings.Index(sentence, "Source")
	fmt.Println(index)
}
