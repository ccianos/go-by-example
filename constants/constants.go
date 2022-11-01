package main

import (
	"fmt"
	"math"
)

const s string = "conststant"

func main() {
	fmt.Println(s)

	const n = 5000000000

	fmt.Println(n)

	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(float64(d))
	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))

}
