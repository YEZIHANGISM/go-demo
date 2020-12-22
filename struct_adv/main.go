package main

import (
	"fmt"
)

// Any empty interface
type Any interface{}

// Car struct
type Car struct {
	Model       string
	Manufacture string
	BuildYear   int
}

// Cars slice of Car
type Cars []*Car

func main() {
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacture == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Printf("allCars: %#v\n", allCars)
	fmt.Printf("allNewBMWs: %#v\n", allNewBMWs)

	manuFacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Porsche"}
	sortedAppender, sortedCars := MakeSortedAppender(manuFacturers)
	allCars.Process(sortedAppender)
	fmt.Printf("Map sortedCars: %#v\n", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Printf("We have %d BMWs\n", BMWCount)
}

// Process factory function
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		//fmt.Println(annoy, c)
		f(c)
	}
}

// FindAll find all sub queryset
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

// MakeSortedAppender sorting cars according manufacturers
func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacture]; ok {
			sortedCars[c.Manufacture] = append(sortedCars[c.Manufacture], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}
