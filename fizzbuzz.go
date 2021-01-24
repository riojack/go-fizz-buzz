package fizzbuzz

import (
	"fmt"
)

func fizzbuzz(number int) string {
	var output = fmt.Sprintf("%v", number)

	if number == 0 {
		output = "0"
	} else if number % 3 == 0 && number % 5 == 0 {
		output = "fizzbuzz"
	} else if number % 3 == 0 {
		output = "fizz"
	} else if number % 5 == 0 {
		output = "buzz"
	}
	
	return output
}
