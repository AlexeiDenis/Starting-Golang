package main

import (
	"fmt"
)

// func adun1(x int, y int) {
// 	var i int = 0
// 	i = x + y
// 	fmt.Println(i, reflect.TypeOf(i))
// 	//cond := reflect.typeof(x)
// }
// func adun2(x float32, y float32) {
// 	fmt.Println(x+y, reflect.TypeOf(x+y))
// }
// func main() {
// 	adun1(2, 3)
// 	x := -5 + 5
// 	adun2(22.12, 33.12)
// 	fmt.Println(x, reflect.TypeOf(x))
// 	text := "Ionut"
// 	text1 := "Jon" + " " + text
// 	fmt.Println(text, reflect.TypeOf(text))
// 	fmt.Println(text1, reflect.TypeOf(text1))
// 	var kkk int = 15136362
// 	rezultat := strconv.Itoa(kkk)
// 	fmt.Println(rezultat, reflect.TypeOf(rezultat))
// }
func main() {
	// var array = [5]int{4, 5, 6, 7, 12}
	//sau
	// array1 := [5]int{4, 5, 6, 7, 12}

	var arraytest = [7]int{1, 7, 3, 4, 5, 6, 2}

	// fmt.Println(array)
	// fmt.Println(array[1] + array[0])

	// fmt.Println(array1, reflect.TypeOf(array1))

	// fmt.Println(array[1] + array1[0])

	// fmt.Println(array[0] + arraytest[6]) //merge adunat doar daca sunt de acelasi tip declarate

	for i := 0; i < len(arraytest); i++ {
		for j := i + 1; j < len(arraytest); j++ {
			if arraytest[i] < arraytest[j] {
				schimb := arraytest[i]
				arraytest[i] = arraytest[j]
				arraytest[j] = schimb
			}

		}

	}
	for i := 0; i < len(arraytest); i++ {
		fmt.Println(arraytest[i])
	}
}
