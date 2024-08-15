package main

import "fmt"

type IntCollection struct {
	data []int
}

func NewIntCollection(data []int) *IntCollection {
	return &IntCollection{data: data}
}

// Iterator returns an iterator function for the collection
func (c *IntCollection) Iterator() func() (int, bool) {
	index := 0
	return func() (int, bool) {
		if index < len(c.data) {
			value := c.data[index]
			index++
			return value, true
		}
		return 0, false // No more items
	}
}

func main() {
	collection := NewIntCollection([]int{1, 2, 3, 4, 5})

	iter := collection.Iterator()

	// Use the iterator in a for-range loop
	for value, ok := iter(); ok; value, ok = iter() {
		fmt.Println(value)
	}
}
