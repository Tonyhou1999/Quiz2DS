package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func computeSum(input []int) int {
	sum := 0
	// We are using a for loop to calculate the sum of all the elements
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum
}

func computeSumNew(input []int, output chan int) {
	sum := 0
	for _, num := range input {
		sum += num
	}
	output <- sum
}
func main() {
	//We are getting an input from the command line, but this is a string, so we will convert it to an int in the next line
	nums := os.Args[1]
	bound, _ := strconv.ParseInt(nums, 0, 32)
	var solution = make([]int, bound+1)
	for i := 0; i < len(solution); i++ {
		//This is to generated random integers and store them in a slice, this fmt.Printf statement below is to assist to check whether it suceessfully generated or not, uncomennt will reveal the prints
		solution[i] = rand.Intn(int(bound))
		//fmt.Printf("The %d index has randomly generated value of %d\n", i, solution[i])
	}
	// First we will see how long it takes to compute the sum directly
	SumBefore := time.Now()
	fmt.Printf("The sum of the array is %d\n", computeSum(solution))
	SumNow := time.Since(SumBefore)
	fmt.Printf("It takes %d nanoseconds to compute the summation directly\n", SumNow.Nanoseconds())

	//Now we will see how long it takes running concurrently, following a similar approach
	Now := time.Now()
	temp1 := make(chan int)
	temp2 := make(chan int)
	go computeSumNew(solution[:len(solution)/2], temp1)
	go computeSumNew(solution[:len(solution)/2], temp2)

	// We will receive the output directly from each subarray
	x, y := <-temp1, <-temp2
	fmt.Printf("The sum of the array is %d\n", x+y)
	End := time.Since(Now)
	fmt.Printf("It takes %d nanoseconds to compute the summation directly\n", End.Nanoseconds())

}
