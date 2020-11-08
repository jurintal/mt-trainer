package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var ErrEmptySlice = errors.New("Ei yhtään kertotaulua valittuna!")

func main() {
	var (
		answer         string                          // User given answer
		totalCorrect   int                             // Total count of correct answers
		totalWrong     int                             // Total count of wrong answers
		targets        = []int{2, 3, 4, 5, 6, 7, 8, 9} // Target values to multiply
		multipliers    []int                           // Selected multipliers
		prevMultiplier int                             // Previous multiplier asked
		prevTarget     int                             // Previous target asked
	)

	// Get valid multipliers from program arguments (=integer values)
	sourceMultipliers := os.Args[1:]
	err, multipliers := GetIntegersFromSlice(&sourceMultipliers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Harjoitellaan kertotauluja: %v. Vastaa 'q' kun haluat lopettaa.\n", multipliers)

	// Ask questions until "q" given
	for {
		err, rm, rt := GetRandomValuesFromSlices(&multipliers, &targets, prevMultiplier, prevTarget)
		if err != nil {
			fmt.Println(err)
			return
		}
		prevMultiplier, prevTarget = rm, rt
		// Correct answer for this round's question
		correct := rm * rt

		fmt.Printf("Paljonko on %d * %d\n", rm, rt)
		// Read the answer
		fmt.Scanln(&answer)
		// If quitting, show statistics
		if answer == "q" {
			fmt.Printf("Sait yhteensä %d OIKEIN ja %d VÄÄRIN\n", totalCorrect, totalWrong)
			return
		}
		// Check the answer - wrong if not integer answer or not matching correct answer
		a, err := strconv.Atoi(answer)
		if err != nil || a != correct {
			totalWrong++
			fmt.Printf("Väärin, oikea vastaus: %d\n", correct)
		}
		// Correct answer
		if a == correct {
			totalCorrect++
			fmt.Print("OIKEIN\n")
		}
	}
}

// Get valid integer multipliers from the slice of stings
func GetIntegersFromSlice(source *[]string) (err error, multipliers []int) {
	for _, arg := range *source {
		m, err := strconv.Atoi(arg)
		if err == nil {
			multipliers = append(multipliers, m)
		}
	}
	// Check that we have at least 1 multiplier given
	if len(multipliers) == 0 {
		err = ErrEmptySlice
		return
	}
	return
}

// Return random values from given slices
func GetRandomValuesFromSlices(multipliers, targets *[]int, skipMultiplier, skipTarget int) (err error, mul, tar int) {
	var i = 0 // For making sure that we do not end up in never ending loop
	// Check that we have values in both slices
	if len(*multipliers) == 0 || len(*targets) == 0 {
		err = ErrEmptySlice
		return
	}
	for {
		i++
		// Seed the random number generator or the sequence will be the same each time
		rand.Seed(time.Now().UnixNano())
		// Get the multiplier and target randomly
		mul = (*multipliers)[rand.Intn(len(*multipliers))]
		tar = (*targets)[rand.Intn(len(*targets))]
		// Exit the loop only when at least one of the randomly selected values has changed (or we have tried 10 times already)
		if mul != skipMultiplier || tar != skipTarget || i > 10 {
			break
		}
	}
	return
}
