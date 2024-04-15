package sugar

import (
	"errors"
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
	expectedValues := []int{4, 5}
	waitDoubleOne := Async(waitAndSum, 2, 2)
	waitDoubleTwo := Async(waitAndSum, 2, 3)
	gots, _ := AwaitAll(waitDoubleOne, waitDoubleTwo)
	for i := 0; i < len(gots); i++ {
		if gots[i] != expectedValues[i] {
			t.Errorf("Expected %d, got %d", expectedValues[i], gots[i])
		}
	}
}

func TestAwaitAllErr(t *testing.T) {
	expectedValues := []int{4, 0}
	expectedErrs := []error{nil, errSimulated}

	waitDouble := Async(waitAndSum, 2, 2)
	waitErr := Async(waitAndErr, 2, 3)
	got, errs := AwaitAll(waitDouble, waitErr)
	for i := 0; i < len(errs); i++ {
		if errs[i] != expectedErrs[i] {
			t.Errorf("Expected %s, got %s", expectedErrs[i], errs[i])
		}
	}
	for i := 0; i < len(got); i++ {
		if got[i] != expectedValues[i] {
			t.Errorf("Expected %d, got %d", expectedValues[i], got[i])
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

func TestAwaitErr(t *testing.T) {
	waitErr := Async(waitAndErr, 1, 2)
	got, err := Await(waitErr)
	if err == nil {
		t.Errorf("Expected an error, got %d", got)
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

var errSimulated = errors.New("some error")

func waitAndErr(nums ...int) (int, error) {
	time.Sleep(2 * time.Second)
	return 0, errSimulated
}
