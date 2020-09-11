package Utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Function that takes console input as string and converts it to int
func Input() (map[int]int, map[int]int) {
Bet:
	fmt.Println("Enter a Bet and Result in format (BetScoreT1/BetScoreT2),(ResultScoreT1/ResultScoreT2)")

	mergedInput := InputLoop()

	bet, err := InputExtractor(mergedInput, 0)
	if err != nil{
		fmt.Println(err)
		goto Bet
	}

	res, err := InputExtractor(mergedInput, 1)
	if err != nil{
		fmt.Println(err)
		goto Bet
	}

	return bet, res
}

// Function that takes console input and check if it's in valid format. If not loops until user enters correct input.
func InputLoop() [][]string {
	var (
		inpt string // users input
		splitedInput []string

		betInput []string
		resInput []string
	)
	// Label for Loop to break it if all correct
InputLoop:
	for{
		fmt.Scanf("%s", &inpt)

		switch splitedInput = strings.Split(inpt, ","); {

		// If splited input by "," less than 2 raises continue looping
		case len(splitedInput) != 2:
			fmt.Println("Wrong input format")

		case len(splitedInput) == 2:

			switch betInput, resInput =
				strings.Split(splitedInput[0], "/"), strings.Split(splitedInput[1], "/");{

			// If both inputs splited by "/" len equals to 2 breaks loop
			case len(betInput) == 2 && len(resInput) == 2:
				break InputLoop

				default:
				fmt.Println("Wrong input format")
			}
		}
	}

	mergedInput := [][]string{
		betInput, resInput,
	}
	return mergedInput
}

//Function to extract results from bet/res by input and returns ExtractedInput and error
func InputExtractor(mergedInput [][]string, inptnum int) (map[int]int, error) {
	var (
		ExtractedInput = make(map[int]int)
		err error
	)

	for i, bt := range mergedInput[inptnum]{
		// Checks if int was input
		if bti, err1 := strconv.Atoi(bt); err1 == nil{
			ExtractedInput[i+1] = bti
		}else {
			err = fmt.Errorf("Wrong input in %v.\n", inptnum)
		}
	}

	return ExtractedInput, err
}