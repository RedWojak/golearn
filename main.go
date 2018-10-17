package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {

	w1, w2 := "Hey", "there"
	var a int = 63

	fmt.Println(multiple(w1, w2))

}
