package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	//Everything is the same as above
	nums := os.Args[1]
	bound, _ := strconv.ParseInt(nums, 0, 32)

	var SliceSort = make([]int, bound)
	var StableSort = make([]int, bound)
	//The line below is to check the type of what we are looking for, just in case we are getting not a slice(that means there is an error)
	//fmt.Println("The type of the array of randomly generated integers, is indeed, a", reflect.TypeOf(solution))

	//Looping through the entire slice, and randomly generated an integer, I set a bound though.
	//I created another copy of the original randomly generated slice, in order to perform the stable sort on it and see the time complexity.
	for i := 0; i < len(SliceSort); i++ {
		SliceSort[i] = rand.Intn(int(bound))
		StableSort[i] = SliceSort[i]
		//fmt.Printf("The %d index has randomly generated value of %d\n", i, solution[i])
	}

	//Now we will sort them, and see how long the program runs.
	start := time.Now()
	sort.Ints(SliceSort)
	duration := time.Since(start)

	StartStable := time.Now()
	sort.SliceStable(StableSort, func(p, q int) bool {
		return StableSort[p] < StableSort[q]
	})
	EndStable := time.Since(StartStable)

	//Now we have computed the running time, so will print them out in statement and see how long does it cost
	fmt.Printf("The Sorting used to sort the slice takes %d nanoseconds to sort\n", duration.Nanoseconds())
	fmt.Printf("The SliceStable used to sort the slice takes %d nanoseconds to sort\n", EndStable.Nanoseconds())
}
