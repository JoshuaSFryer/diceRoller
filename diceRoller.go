// diceRoller prompts the user to input a dice roll in the format NdM, where
// N represents the number of dice to roll, each of which having M sides.
// d is used to separate the numbers.
// For example, 3d6 translates to a roll of 3, 6-sided dice.
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
)

func main() {
	seed()
	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt the user and save the input.
		fmt.Print("Enter a dice roll: ")
		text, _ := reader.ReadString('\n')

		// Split the input and separate the two numbers, storing them in vars.
		roll := strings.Split(text, "d")
		number, _ := strconv.Atoi(roll[0])
		sides, _ := strconv.Atoi(strings.Trim(roll[1], "\n"))

		// Roll the specified dice and print the result.
		fmt.Printf("Rolling %d dice of %d sides. Results: ", number, sides)
		results, _ := rollNDice(number, sides)
		sum := 0
		for _, num := range results {
			if num != 0 {
				fmt.Printf("%d ", num)
				sum += num
			}
		}
		fmt.Printf("\n Total: %d\n", sum)
	}
}

// rollDie rolls a single die, of the specified number of sides.
// Side number can range from 2 to 100, inclusive.
// It returns the result of the roll, and any error encountered.
func rollDie(sides int) (int, error) {
	// Validate input.
	if sides < 2 || sides > 100 {
		err := errors.New("Number of sides must be between 2-100")
		fmt.Println(err)
		return -1, err
	}
	// rand.Intn returns a number between 0 and n-1.
	// Add 1 to the result to get between 1 and n, as with an n-sided die.
	return rand.Intn(sides) + 1, nil
}

// rollNDice rolls a specified number of dice, of an arbitrary number of sides.
// It returns a slice containing the results of each roll, and any error
// encountered.
func rollNDice(num int, sides int) ([]int, error) {
	// Validate input.
	if num < 1 {
		err := errors.New("Must roll at least 1 die")
		fmt.Println(err)
		return make([]int, 0), err
	}

	// Roll the dice and sum up the results.
	// Create a slice with capacity num.
	rolls := make([]int, num)
	for i := 0; i < num; i++ {
		r, err := rollDie(sides)
		if err == nil {
			// Add the roll to the list of rolls.
			rolls = append(rolls, r)
		} else {
			// Pass on the error thrown by RollDie
			return make([]int, 0), err
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
