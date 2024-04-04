# Sugar
### Overview
Sugar is a go library that tries to provide js like `async/await` construct using goroutine and channel

### Installing
```cURL
go get -u github.com/shafn098/sugar@latest
```

### Usage
```go
package main

import (
	"fmt"
	"time"
	"github.com/shafin098/sugar"
)

// WaitTwoSecondsAndSum simulates long function call
func WaitTwoSecondsAndSum(nums ...int) (int, error) {
	time.Sleep(2 * time.Second)
	return nums[0] + nums[1], nil
}

// WaitFourSecondsAndSum simulates long function call
func WaitFourSecondsAndSum(nums ...int) (int, error) {
	time.Sleep(4 * time.Second)
	return nums[0] + nums[1], nil
}

func main() {
    // sum one will be running in backgroung in a separate goroutine without blocking until await is called
    waitSumOne := sugar.Async(WaitTwoSecondsAndSum, 2, 3)
    fmt.Println("summing 2,3 and waiting for 2 seconds")

    // sum two will be running in backgroung in a separate goroutine without blocking until await is called
    waitSumTwo, err := sugar.Async(WaitTwoSecondsAndSum, 3, 3)
    fmt.Println("summing 3,3 and waiting for 4 seconds")

    // we are supposed to do other work here
    // -------------------------------------
    // -------------------------------------
    // -------------------------------------
    // assuming other works are done

    // blocking main goroutine and fetching first sum, if it has been more than 2 seconds (other work took 2 seconds) await call will return instantly with sum
    sumOne, err := sugar.Await(waitSumOne)
    if err != nil {
        fmt.Errorf(err)
    }
    fmt.Println("sumOne: ", sumOne)


    // blocking main goroutine and fetching second sum, if it has been more than 4 seconds (other work took 4 seconds) await call will return instantly with sum
    sumTwo, err := sugar.Await(waitSumTwo)
    if err != nil {
        fmt.Errorf(err)
    }
    fmt.Println("sumTwo: ", sumTwo)
}
```
