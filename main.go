package main

import "fmt"

func main() {
	tossesTotal := 0	// total of tosses for total computation

	die1 := defaultDie()	// create new die
	die2 := copyDie(&die1)	// create second die copying the first one

	for i := 1; i < 4; i++ {
		die1.roll()
		die2.roll()
		tossValue := die1.getFaceValue() + die2.getFaceValue()	// value of one toss

		fmt.Printf("Toss %v generated a %v and a %v for a total of\t:\t %v\n",
			i, die1.getFaceValue(), die2.getFaceValue(), tossValue)

		tossesTotal += tossValue
	}
	fmt.Printf("The total of tosses is\t\t\t\t:\t %v\n", tossesTotal)
}
