package main

import (
	"fmt"
	"go-design-patterns/structural/decorator"
)

func main() {
	fmt.Println("tax on 28L income: ", decorator.computeTax(400000))
}
