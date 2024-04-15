package main

import (
	"fmt"
	"os"
	"time"

	sugar "github.com/shafin098/sugar"
)

func WaitTwoSecondsAndSum(nums ...int) (int, error) {
	time.Sleep(2 * time.Second)
	return nums[0] + nums[1], nil
}

func WaitFourSecondsAndSum(nums ...int) (int, error) {
	time.Sleep(4 * time.Second)
	return nums[0] + nums[1], nil
}

func main() {
	waitSumOne := sugar.Async(WaitTwoSecondsAndSum, 2, 3)
	fmt.Println("summing 2,3 and waiting for 2 seconds")

	waitSumTwo := sugar.Async(WaitTwoSecondsAndSum, 3, 3)
	fmt.Println("summing 3,3 and waiting for 4 seconds")

	sumOne, err := sugar.Await(waitSumOne)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
	}
	fmt.Println("sumOne: ", sumOne)

	sumTwo, err := sugar.Await(waitSumTwo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
	}
	fmt.Println("sumTwo: ", sumTwo)
}
