package assignments

import "fmt"

func CreateMap() {
	mappy := map[string]int{"one": 1, "two": 2, "five": 5}
	fmt.Println(mappy)
	DeleteMap(mappy)
	fmt.Println(mappy)
}

func DeleteMap(mapArg map[string]int) {
	for key, _ := range mapArg {
		delete(mapArg, key)
	}
}
