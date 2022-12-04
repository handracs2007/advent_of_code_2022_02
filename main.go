package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// in checkes whether an element exists in the slice of elements. Returns true if it is, otherwise false.
func in(elements []uint8, element uint8) bool {
	for i := 0; i < len(elements); i++ {
		if elements[i] == element {
			return true
		}
	}

	return false
}

func part1() {
	// Read the input file.
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %s", err)
	}
	defer f.Close()

	// Indication for the rock, paper, and scissor.
	rock, paper, scissor := []uint8{uint8('A'), uint8('X')}, []uint8{uint8('B'), uint8('Y')}, []uint8{uint8('C'), uint8('Z')}

	// Creates a map of score for each element.
	scores := make(map[uint8]int)
	scores[rock[1]] = 1
	scores[paper[1]] = 2
	scores[scissor[1]] = 3

	// Additional score for each situation.
	lost, draw, win := 0, 3, 6

	// An initial score of 0.
	score := 0
	r := bufio.NewReader(f)
	for {
		l, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("error whilst reading input file: %s", err)
		}

		g := strings.Split(l, " ")
		o, i := g[0][0], g[1][0] // First element o is what is thrown by other, and second element i is what we have.
		score += scores[i]       // Add the score with the element that we throw.
		if in(rock, o) {
			if in(rock, i) {
				score += draw
			} else if in(paper, i) {
				score += win
			} else {
				score += lost
			}
		} else if in(paper, o) {
			if in(paper, i) {
				score += draw
			} else if in(rock, i) {
				score += lost
			} else {
				score += win
			}
		} else if in(scissor, o) {
			if in(scissor, i) {
				score += draw
			} else if in(rock, i) {
				score += win
			} else {
				score += lost
			}
		}

		if err == io.EOF {
			break
		}
	}

	fmt.Println(score)
}

func part2() {
	// Read the input file.
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %s", err)
	}
	defer f.Close()

	// Character indication for each element.
	rock, paper, scissor := uint8('A'), uint8('B'), uint8('C')

	// Character indication for each target situation.
	lose, draw, win := uint8('X'), uint8('Y'), uint8('Z')

	// Additional score we get for each situation.
	loses, draws, wins := 0, 3, 6

	// Creates a map of score for each element.
	scores := make(map[uint8]int)
	scores[rock] = 1
	scores[paper] = 2
	scores[scissor] = 3

	score := 0
	r := bufio.NewReader(f)
	for {
		l, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("error whilst reading input file: %s", err)
		}

		g := strings.Split(l, " ")
		o, t := g[0][0], g[1][0] // First element o is what the other throws, and second element t is the target situation.
		if t == lose {
			score += loses

			switch o {
			case rock:
				score += scores[scissor]
			case paper:
				score += scores[rock]
			case scissor:
				score += scores[paper]
			}
		} else if t == draw {
			score += draws
			score += scores[o]
		} else if t == win {
			score += wins

			switch o {
			case rock:
				score += scores[paper]
			case paper:
				score += scores[scissor]
			case scissor:
				score += scores[rock]
			}
		}

		if err == io.EOF {
			break
		}
	}

	fmt.Println(score)
}

func main() {
	part1()
	part2()
}
