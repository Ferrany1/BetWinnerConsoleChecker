package Betting

// Winner type constants
const(
	BigWinner = 2
	SmallWinner = 1
	Looser = 0
)

// Function that decides prize for bet and returns int eq
func Decider(bet map[int]int, res map[int]int) int {
	var(
		BetIsDraw = bet[1] == bet[2]
		BetFTWin = bet[1] > bet[2]

		ResIsDraw = res[1] == res[2]
		ResFTWin = res[1] > res[2]

		result int
		// True False
		// False False

		// False False
		// False False
	)
	switch {
	// Checking if bet wins
	case (BetFTWin == ResFTWin && ((BetIsDraw == ResIsDraw)	||
		(BetIsDraw != ResFTWin && BetFTWin != ResIsDraw))) ||
		(BetIsDraw == ResIsDraw && BetIsDraw != false && ResIsDraw != false):
		// Checking for each bet/res element
		for i := 0; i <= len(bet); i++ {
			// BigWinner if bet and result are equal
			if bet[i] == res[i]{
				result = BigWinner
			// BigWinner if bet and result are in the same direction
			}else{
				result = SmallWinner
				break
			}
		}
	default:
		// Looser in all other cases
		result = Looser
	}

	return result
}