package main

import "fmt"

// Car Struct
type Car struct {
	year int
	make *string
}

func (c Car) String() string {
	return fmt.Sprintf("{make: %s, year: %d}", *c.make, c.year)
}

func main() {
	make := "Toyota"
	myCar := Car{year: 1996, make: &make}
	fmt.Println(myCar)
}
