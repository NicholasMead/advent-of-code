// Declaration of the main package
package main

// Importing packages
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// Main function
func main() {
	start := time.Now().UnixMicro()

	instructions := readInput(os.Args[1])
	round1Scores := make(chan uint, 15)
	round2Scores := make(chan uint, 15)

	var wg sync.WaitGroup

	for n := 0; n < 5; n++ {
		wg.Add(1)
		go func() {
			count := 0
			for instruction := range instructions {
				round1Scores <- calculateRound1Score((instruction))
				round2Scores <- calculateRound2Score((instruction))
				count++
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(round1Scores)
		close(round2Scores)
	}()

	round1 := sum(round1Scores)
	round2 := sum(round2Scores)

	fmt.Println("Round 1, score:", <-round1)
	fmt.Println("Round 2, score:", <-round2)

	end := time.Now().UnixMicro()
	fmt.Println("Execution time (us):", end-start)

}

func sum(c <-chan uint) <-chan uint {
	output := make(chan uint)

	go func() {
		var sum uint = 0
		for value := range c {
			sum += value
		}
		output <- sum
		defer close(output)
	}()

	return output
}

func readInput(path string) <-chan string {
	var stream *os.File

	switch path {
	case "":
		panic("input path needed")
	case "-":
		stream = os.Stdin
	default:
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		stream = file
	}

	reader := bufio.NewReader(stream)
	output := make(chan string, 2500)

	go func() {
		for {
			input, _ := reader.ReadString('\n')
			input = strings.Trim(input, "\n ")
			if input != "" {
				output <- input
			} else {
				close(output)
				return
			}
		}
	}()

	return output
}

func calculateRound1Score(instruction string) uint {
	codedMoves := strings.Split(instruction, " ")
	opponentMove := getMoveValue(codedMoves[0])
	playerMove := getMoveValue(codedMoves[1])

	var score uint = 0

	//draw
	if playerMove == opponentMove {
		score += 3 //draw
	} else if playerMove == ((opponentMove + 1) % 3) {
		score += 6 //win
	} else {
		score += 0 //loss
	}

	score += uint(playerMove + 1)

	return score
}

func calculateRound2Score(instruction string) uint {
	codedMoves := strings.Split(instruction, " ")
	opponentMove := getMoveValue(codedMoves[0])
	playerAction := codedMoves[1]
	var playerMove uint8

	var score uint = 0

	switch playerAction {
	case "X": //lose
		score += 0                          //for loosing
		playerMove = (opponentMove + 2) % 3 //for move
	case "Y": //draw
		score += 3                //for drawing
		playerMove = opponentMove //for move
	case "Z": //win
		score += 6                          //for winning
		playerMove = (opponentMove + 1) % 3 //for move
	default:
		panic("Unknown plan!")
	}

	score += uint(playerMove + 1)

	return score
}

func getMoveValue(code string) uint8 {
	switch code {
	case "A":
		return 0
	case "B":
		return 1
	case "C":
		return 2
	case "X":
		return 0
	case "Y":
		return 1
	case "Z":
		return 2
	default:
		return 0
	}
}
