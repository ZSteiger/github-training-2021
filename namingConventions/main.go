package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type fancyShapeArithmetic interface {
	insidebits() float64
	outerbounds() float64
}

type multi_sided_quadrilateral struct {
	theSides, howTallAmI float64
}

type idk struct {
	halfThe_diameter float64
}

func (instantiated_multi_sided_quadrilateral multi_sided_quadrilateral) insidebits() float64 {
	return instantiated_multi_sided_quadrilateral.theSides * instantiated_multi_sided_quadrilateral.howTallAmI
}

func (shape multi_sided_quadrilateral) outerbounds() float64 {
	return 2*shape.theSides + 2*shape.howTallAmI
}

func (fuckCircles idk) insidebits() float64 {
	return math.Pi * fuckCircles.halfThe_diameter * fuckCircles.halfThe_diameter
}

func (what idk) outerbounds() float64 {
	return 2 * math.Pi * what.halfThe_diameter
}

func takeTheInternalMeasurementsOfTheItemInQuestionAndReportBackTheResultOfTheBasicOperandsNecessaryForDeterminingTheCode(asdfghjkl fancyShapeArithmetic) {
	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("Today is the day 😎")
	default:
		fmt.Println("Keep dreaming 😭")
	}

	fmt.Printf("Result: %v", asdfghjkl)
	fmt.Printf("Area: %v", asdfghjkl.insidebits())
	fmt.Printf("Perimeter: %v", asdfghjkl.outerbounds())
}

func sundayChores(hamper []string) {
	sort.Strings(hamper)
	fmt.Println(hamper)
}

// Go programs require package main, and func main to be declared as an entry point
// no need to rename them
func main() {

	msq := multi_sided_quadrilateral{theSides: 3, howTallAmI: 4}
	istilldontknow := idk{halfThe_diameter: 5}

	takeTheInternalMeasurementsOfTheItemInQuestionAndReportBackTheResultOfTheBasicOperandsNecessaryForDeterminingTheCode(msq)
	takeTheInternalMeasurementsOfTheItemInQuestionAndReportBackTheResultOfTheBasicOperandsNecessaryForDeterminingTheCode(istilldontknow)

	dirty_laundry := []string{"underwear", "socks", "old tshirts", "work shorts", "facemask"}
	sundayChores(dirty_laundry)
}
