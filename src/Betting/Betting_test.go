package Betting

import (
	"testing"
)

type TestCases struct {
	TestBet map[int]int
	TestRes map[int]int
	Result int
}

var Tests = []TestCases{
	{TestBet: map[int]int{1:1, 2:2}, TestRes: map[int]int{1:3, 2:4}, Result: 1}, // If direction of prediction was correct
	{TestBet: map[int]int{1:1, 2:1}, TestRes: map[int]int{1:1, 2:1}, Result: 2}, // If Score is equal
	{TestBet: map[int]int{1:1, 2:1}, TestRes: map[int]int{1:2, 2:2}, Result: 1}, // If score is not equal and prediction wasn't decided
	{TestBet: map[int]int{1:1, 2:1}, TestRes: map[int]int{1:3, 2:4}, Result: 0}, // If score is not equal and prediction wasn't decided
	{TestBet: map[int]int{1:1, 2:2}, TestRes: map[int]int{1:4, 2:3}, Result: 0}, // If Direction was wrong
}

func TestDecider(t *testing.T) {
	for ind, test := range Tests{
		if res := Decider(test.TestBet, test.TestRes); res != test.Result{
			t.Errorf("Test %v error. Result is %v, want %v", ind + 1, res, test.Result)
		}
	}
}
