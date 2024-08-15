package main

import (
	"fmt"
	"maps"
	"slices"
	"sync"
)

func main() {
	//Range iterator using concurrent safe map
	fmt.Printf("--> Range iterator using concurrent safe map\n")
	var syncMap sync.Map

	syncMap.Store("k1", "v1")
	syncMap.Store("k2", "v2")
	syncMap.Store("k3", "v3")

	for key, value := range syncMap.Range {
		fmt.Println(key, value)
	}

	//Slice iterator
	fmt.Printf("\n--> Slice iterator sorted\n")
	slice := []string{"d", "a", "e", "b", "z", "c"}
	sortedSlice := slices.Sorted(slices.Values(slice))
	for key, value := range slices.All(sortedSlice) {
		fmt.Printf("%d:%v ", key, value)
	}

	//Map iterator
	fmt.Printf("\n\n--> Map iterator\n")
	mapIterator := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for key, value := range maps.All(mapIterator) {
		fmt.Printf("%v:%v ", key, value)
	}
}
