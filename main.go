package main

import (
	"fmt"
	"stateroutine/stateroutine"
)

// Example usage of stateroutine
func main() {
	state := stateroutine.Go[int]()

	stateroutine.Set(state, "one", 1)
	stateroutine.Set(state, "two", 2)

	value := stateroutine.Get(state, "one")
	if value != nil {
		fmt.Println(*value)
	}
	value = stateroutine.Get(state, "two")
	if value != nil {
		fmt.Println(*value)
	}

	value = stateroutine.Get(state, "three")
	if value != nil {
		fmt.Println(*value)
	}

	stateroutine.Delete(state, "one")
	value = stateroutine.Get(state, "one")
	if value != nil {
		fmt.Println(*value)
	}
	value = stateroutine.Get(state, "two")
	if value != nil {
		fmt.Println(*value)
	}
}
