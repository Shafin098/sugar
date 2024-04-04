# Sugar
Sugar is a go library that tries to provide js like async/await construct using goroutines and channel

# Usage example
```go
package main

import (
	"fmt"
	"time"
	"github.com/shafin098/sugar"
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

	waitSumTwo, err := sugar.Async(WaitTwoSecondsAndSum, 3, 3)
	fmt.Println("summing 3,3 and waiting for 4 seconds")

	sumOne, err := sugar.Await(waitSumOne)
	if err != nil {
		fmt.Errorf(err)
	}
	fmt.Println("sumOne: ", sumOne)

	sumTwo, err := sugar.Await(waitSumTwo)
	if err != nil {
		fmt.Errorf(err)
	}
	fmt.Println("sumTwo: ", sumTwo)
}

```