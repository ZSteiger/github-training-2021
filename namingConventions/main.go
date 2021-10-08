package main

import (
	"fmt"
	"math"
	"sort"
)

type geometricCalculations interface {
	getArea() float64
	getPerimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (rect Rectangle) getArea() float64 {
	return rect.width * rect.height
}

func (shape Rectangle) getPerimeter() float64 {
	return 2*shape.width + 2*shape.height
}

func (circ Circle) getArea() float64 {
	return math.Pi * circ.radius * circ.radius
}

func (circ Circle) getPerimeter() float64 {
	return 2 * math.Pi * circ.radius
}

func printAreaAndPerimeter(shape geometricCalculations) {
	// TODO: Delete commented code below b/c why 

	// switch time.Now().Weekday() {
	// case time.Friday:
	// 	fmt.Println("Today is the day 😎")
	// default:
	// 	fmt.Println("Keep dreaming 😭")
	// }

	fmt.Printf("Result: %v, ", shape)
	fmt.Printf("Area: %v, ", shape.getArea())
	fmt.Printf("Perimeter: %v \n", shape.getPerimeter())
}

func printSortedStrings(array []string) {
	sort.Strings(array)
	fmt.Println(array)
}

// Go programs require package main, and func main to be declared as an entry point
// no need to rename them
func main() {

	rect := Rectangle{width: 3, height: 4}
	circ := Circle{radius: 5}

	printAreaAndPerimeter(rect)
	printAreaAndPerimeter(circ)

	dirty_laundry := []string{"underwear", "socks", "old tshirts", "work shorts", "facemask"}
	printSortedStrings(dirty_laundry)
}
