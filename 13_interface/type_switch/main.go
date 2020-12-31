package main

import (
	"fmt"
	"strings"
)

func main() {
	explain("Hello World")
	explain(23)
	explain(true)
}
func explain(i interface{}) {
	// Type switch
	switch i.(type) {
	case string:
		fmt.Println("i stored string", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}
