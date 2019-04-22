package example

import (
	"bufio"
	"fmt"
	"go-learn/cmd"
	"math"
	"math/rand"
	"os"
	"strings"
)

const percentPerStep float64 = 0.5
const winningScore int = 10

func StartPanicAttack() {

	cmd.Println("Panic Attack The Game. Press 'h' to get help")

	// If nil then quit
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("You lose. Your score is ", r)
		}
	}()

	step(0)
}

func step(score int) {
	if score == winningScore {
		fmt.Println("You win!!! Your score is ", score)
		panic(nil)
	}

	fmt.Println("Your score is ", score)

	reader := bufio.NewReader(os.Stdin)

	for {
		command := cmd.ReadString(reader)

		switch {
		case strings.Compare(command, "a") == 0:
			if isDefeated(score) {
				panic(score)
			}

			step(score + 1)
			return
			break

		case strings.Compare(command, "d") == 0:
			if isDefeated(score) {
				panic(score)
			}

			score--
			defer func() {
				if r := recover(); r == nil {
					panic(nil)
				}
				fmt.Println("You defended with score ", score)
				step(score)
				return
			}()

			step(score)
			return
			break

		case strings.Compare(command, "q") == 0:
			panic(nil)
			break

		case strings.Compare(command, "h") == 0:
			fmt.Println("There is no help!")
			break

		default:
			cmd.Println("You can use only 'a', 'd' and 'q' to quit")
			break
		}
	}
}

func isDefeated(score int) bool {
	return rand.Intn(100) < int(math.Round(float64(score)/percentPerStep))
}
