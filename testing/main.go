package testing

import "fmt"

func main() {
	// initialize a variable containing an int of 1
	starterInt := 1

	result := toTen(starterInt)
	// pass to the function to be incremented
	fmt.Println(result)
}

func toTen(i int) int {
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	return i
}
