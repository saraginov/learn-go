package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	returnMap := make(map[string]int)

	// https://golang.org/pkg/strings/#Fields
	arrayOfStrings := strings.Fields(s) 

	for i := 0; i<len(arrayOfStrings); i++ {
		currentWordWeAreTesting := arrayOfStrings[i]
		
		elem, ok := returnMap[currentWordWeAreTesting]

		if ok {
			/*
				if ok == true, it means word has come up already, thus we increment elem by 1 
				IMPORTANT:
					++ and -- in Go are statements, not expressions
					therefore returnMap[currentWordWeAreTesting] = elem++ returns a syntax error!!
			*/
			elem++
			returnMap[currentWordWeAreTesting] = elem
		} else {
			/* we have not come accross this word so we will add it to our map 
			and initialize its count to 1 */
			returnMap[currentWordWeAreTesting] = 1
		}
	}

	return returnMap
}

func main() {
	wc.Test(WordCount)
}
