package main

import (
	"Test/src/Betting"
	"Test/src/Utils"
	"fmt"
)

//App to get console input of match score bet and actual result and check user Winning
func main()  {
	res := Betting.Decider(Utils.Input())
	fmt.Println(res)
}