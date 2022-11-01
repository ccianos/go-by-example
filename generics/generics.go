package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {}

func main() {
	fmt.Println()
}
