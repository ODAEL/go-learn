package tutorial

import (
	"strconv"
	"studyGo/cmd"
)

const quitCommand rune = 'q'
const examplesListCommand rune = 'l'

func printListHint() {
	text := "Print the number and press Enter to select. You can always print this list with command "
	text += strconv.QuoteRune(examplesListCommand)
	text += " and quite with "
	text += strconv.QuoteRune(quitCommand)
	cmd.Println(text)
}
