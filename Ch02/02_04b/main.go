package main

import "fmt"

func main() {
	i1, i2, i3 := 1, 2, 3
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 2.3, 2.44, 3.44
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)
}
