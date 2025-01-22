package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aValue := 42

	fmt.Println(str1, str2, str3)
	stringLength, error := fmt.Println("The value is ", aValue)
	if error == nil {
		fmt.Println(stringLength)
	}
	fmt.Printf("The string lenght is %d\n", stringLength)
	fmt.Printf("The value has a type of %T\n", stringLength)
}
