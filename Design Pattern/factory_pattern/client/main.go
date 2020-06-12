package main

import (
	"factory/factory"
	"fmt"
	"reflect"
)

func main() {
	car := factory.GetEngine("train")

	fmt.Println(reflect.TypeOf(&car).Elem())

	z := car.Accelerate("car euy")
	train := factory.GetEngine("train")
	y := train.Accelerate("train")
	fmt.Println(z, y)
}
