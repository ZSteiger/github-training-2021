package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type quadrilateral struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (shape quadrilateral) area() float64 {
	return shape.width * shape.height
}

func (shape quadrilateral) perimeter() float64 {
	return 2*shape.width + 2*shape.height
}

func (circ circle) area() float64 {
	return math.Pi * circ.radius * circ.radius
}

func (circ circle) perimeter() float64 {
	return 2 * math.Pi * circ.radius
}

func takeTheInternalMeasurementsOfTheItemInQuestionAndReportBackTheResultOfTheBasicOperandsNecessaryForDeterminingTheCode(asdfghjkl geometry) {
	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("Today is the day 😎")
	default:
		fmt.Println("Keep dreaming 😭")
	}

	fmt.Printf("Result: %v", asdfghjkl)
	fmt.Printf("Area: %v", asdfghjkl.area())
	fmt.Printf("Perimeter: %v", asdfghjkl.perimeter())
}

func sundayChores(hamper []string) {
	sort.Strings(hamper)
	fmt.Println(hamper)
}

// Go programs require package main, and func main to be declared as an entry point
// no need to rename them
func main() {

	area := quadrilateral{width: 3, height: 4}
	perimeter := circle{radius: 5}

	takeTheInternalMeasurementsOfTheItemInQuestionAndReportBackTheResultOfTheBasicOperandsNecessaryForDeterminingTheCode(area)
	takeTheInternalMeasurementsOfTheItemInQuestionAndReportBackTheResultOfTheBasicOperandsNecessaryForDeterminingTheCode(perimeter)

	dirtyLaundry := []string{"underwear", "socks", "old tshirts", "work shorts", "facemask"}
	sundayChores(dirtyLaundry)
}
