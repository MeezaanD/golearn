package maps

import "fmt"

func LearnMaps() {
	// mappy := make(map[int]string)
	// mappy[5] = "five"
	// mappy[1] = "one"

	// mappy := map[string]int{
	// 	"one":  1,
	// 	"five": 5}
	// fmt.Println(mappy)

	mappy := make(map[string]int)
	var ok bool

	mappy["two"] = 0
	delete(mappy, "two")
	_, ok = mappy["two"]
	fmt.Println(ok)
}
