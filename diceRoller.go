// diceRoller prompts the user to input a dice roll in the format NdM, where
// N represents the number of dice to roll, each of which having M sides.
// d is used to separate the numbers.
// For example, 3d6 translates to a roll of 3, 6-sided dice.
// It then prints out the results of each individual roll, as well as the sum
// of the rolls, and their average value.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"diceRoller/roll"
)

const (
	CritFailure int = iota - 1
	CritNeutral
	CritSuccess
)

func main() {
	seed()
	reader := bufio.NewReader(os.Stdin)
	
	for {
		// Prompt the user and save the input.
		fmt.Print("Enter a dice roll of the format XdY: ")
		text, _ := reader.ReadString('\n')

		// Split the input and separate the two numbers, storing them in vars.
		rollStrings := strings.Split(text, "d")
		number, _ := strconv.Atoi(rollStrings[0])
		sides, _ := strconv.Atoi(strings.Trim(rollStrings[1], "\n"))

		// Roll the specified dice and print the results.

		results, err := rollNDice(number, sides)

		if err == nil {
			fmt.Printf("Rolling %d dice of %d sides.\nResults: ", number, sides)
			sum := 0
			for _, r := range results {
				// Ignore the empty slice elements, which initialize to 0.
				if r.Value != 0 {
					r.Print()
					sum += r.Value
				}
			}

			// Print information about the rolls.
			average := (float32(sum) / float32(number))
			fmt.Printf("\nSum: %d\n", sum)
			fmt.Printf("Average: %.3f\n", average)
		}
	}
}

// rollDie rolls a single die, of the specified number of sides.
// Side number can range from 2 to 100, inclusive.
// It returns the result of the roll, whether the roll was a critical success,
// critical failure, or neither, and any error encountered.
func rollDie(sides int) (roll.Roll, error) {
	// Validate input.
	if sides < 2 || sides > 100 {
		err := errors.New("Number of sides must be between 2-100")
		fmt.Println(err)
		return roll.New(0, 0), err
	}

	// rand.Intn returns a number between 0 and n-1.
	// Add 1 to the result to get between 1 and n, as with an n-sided die.
	r := rand.Intn(sides) + 1

	// Rolling the maximum value on a die is a "critical success".
	// Rolling a 1 is a "critical failure".
	crit := CritNeutral

	if r == sides {
		crit = CritSuccess
	} else if r == 1 {
		crit = CritFailure
	}
	return roll.New(r, crit), nil
}

// rollNDice rolls a specified number of dice, of an arbitrary number of sides.
// It returns a slice containing the results of each roll, and any error
// encountered.
func rollNDice(num int, sides int) ([]roll.Roll, error) {
	// Validate num. Validating sides is already handled by rollDie, and we
	// simply pass that error on if it occurs.
	if num < 1 {
		err := errors.New("Must roll at least 1 die")
		fmt.Println(err)
		return nil, err
	}

	// Roll the dice and sum up the results.
	// Create a slice with capacity num.
	rolls := make([]roll.Roll, num)
	for i := 0; i < num; i++ {
		r, err := rollDie(sides)
		if err == nil {
			// Add the roll to the list of rolls.
			rolls = append(rolls, r)
		} else {
			// Pass on the error thrown by RollDie
			return nil, err
		}
	}
	return rolls, nil
}

// seed seeds the PRNG built into the rand library with the current time since
// the Unix epoch.
func seed() {
	// Get time since epoch, in nanoseconds.
	currTime := time.Now().UnixNano()
	rand.Seed(currTime)
}
