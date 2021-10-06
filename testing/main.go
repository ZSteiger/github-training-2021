package main

import "fmt"

func main() {
	// initialize a variable containing an int of 1
	starterInt := 1

	// pass to the function to be incremented
	result := toTen(starterInt)

	// print the result
	fmt.Println(result)
}

func toTen(starterInt int) int {
	// take in variable starterInt,
	// if i is less than 10
	for starterInt <= 10 {
		// print its current value
		fmt.Println(starterInt)
		// increment it by one then reassign it the startInt variable
		starterInt = starterInt + 1
	}
	// return it from the function to be used in main()
	return starterInt
}
