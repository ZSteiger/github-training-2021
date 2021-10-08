package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type shapeInterface interface {
	getArea() float64
	getPerimeter() float64
}

type quadrilateral struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (varQuadrilateral quadrilateral) getArea() float64 {
	return varQuadrilateral.width * varQuadrilateral.height
}

func (shape quadrilateral) getPerimeter() float64 {
	return 2*shape.width + 2*shape.height
}

func (varCircle circle) getArea() float64 {
	return math.Pi * varCircle.radius * varCircle.radius
}

func (varCircle circle) getPerimeter() float64 {
	return 2 * math.Pi * varCircle.radius
}

func printshapeInterface(varShape shapeInterface) {
	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("Today is the day 😎")
	default:
		fmt.Println("Keep dreaming 😭")
	}

	fmt.Printf("Result: %v", varShape)
	fmt.Printf("getArea: %v", varShape.getArea())
	fmt.Printf("getPerimeter: %v", varShape.getPerimeter())
}

func printChores(hamper []string) {
	sort.Strings(hamper)
	fmt.Println(hamper)
}

// Go programs require package main, and func main to be declared as an entry point
// no need to rename them
func main() {

	outRectangle := quadrilateral{width: 3, height: 4}
	outCircle := circle{radius: 5}

	printshapeInterface(outRectangle)
	printshapeInterface(outCircle)

	dirty_laundry := []string{"underwear", "socks", "old tshirts", "work shorts", "facemask"}
	printChores(dirty_laundry)
}
