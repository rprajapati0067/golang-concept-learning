package main

import "fmt"

func main() {

	// Learning behaviour of copy()

	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}

	fmt.Println("a6: ", a6)
	fmt.Println("a4: ", a4)

	copy(a6, a4)
	fmt.Println("a6: ", a6)
	fmt.Println("a4: ", a4)

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6: ", b6)
	fmt.Println("b4: ", b4)

	copy(b4, b6)
	fmt.Println("b6: ", b6)
	fmt.Println("b4: ", b4)

	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}

	copy(s6, array4[1:])
	fmt.Println("array4: : ", array4[1:])
	fmt.Println("s6: ", s6)

	array5 := [4]int{4, -4, 4, -4}
	s7 := []int{1, 1, -1, -1, 5, -5}

	//	copy(array5, s7) // will give compile ime error
	copy(array5[0:], s7)
	fmt.Println("array5: : ", array5[1:])
	fmt.Println("s6: ", s7)

}
