package sugar

import (
	"testing"
	"time"
)

func TestAwaitInt(t *testing.T) {
	expect := 4
	waitDouble := Async(waitReturnDouble, 2)
	got, _ := Await(waitDouble)
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestAwaitIntMultiArgs(t *testing.T) {
	expect := 5
	waitDouble := Async(waitAndSum, 2, 3)
	got, _ := Await(waitDouble)
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestAwaitAll(t *testing.T) {
	expects := []int{4, 5}
	waitDoubleOne := Async(waitAndSum, 2, 2)
	waitDoubleTwo := Async(waitAndSum, 2, 3)
	gots, _ := AwaitAll(waitDoubleOne, waitDoubleTwo)
	for i := 0; i < len(gots); i++ {
		if gots[i] != expects[i] {
			t.Errorf("Expected %d, got %d", expects[i], gots[i])
		}
	}
}

func TestAwaitSring(t *testing.T) {
	expect := "test test"
	waitDouble := Async(waitReturnDoubleString, "test")
	got, _ := Await(waitDouble)
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}
}

func waitReturnDouble(num ...int) (int, error) {
	time.Sleep(2 * time.Second)
	return num[0] * 2, nil
}

func waitReturnDoubleString(s ...string) (string, error) {
	time.Sleep(2 * time.Second)
	return s[0] + " " + s[0], nil
}

func waitAndSum(nums ...int) (int, error) {
	time.Sleep(2 * time.Second)
	return nums[0] + nums[1], nil
}
