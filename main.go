package main

import (
	"fmt"
	"reflect"
)

func adun1(x int, y int) {
	i := 0
	i = x + y
	fmt.Println(i)
	//cond := reflect.typeof(x)
}
func adun2(x float32, y float32) {
	fmt.Println(x + y)
}
func main() {
	adun1(2, 3)
	x := -5 + 5
	adun2(22.12, 33.12)
	fmt.Println(x, reflect.TypeOf(x))
}
