package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzzShouldReturnFizzForNumbersDivisibleByThree(t *testing.T) {
	result := fizzbuzz(3)

	assert.Equal(t, "fizz", result)
}
func TestFizzbuzzShouldReturnBuzzForNumbersDivisibleByFive(t *testing.T) {
	result := fizzbuzz(5)

	assert.Equal(t, "buzz", result)
}

func TestFizzbuzzShouldReturnNumberWhenInputIsNotDivisibleByThreeOrFive(t *testing.T) {
	result := fizzbuzz(1)

	assert.Equal(t, "1", result)
}

func TestFizzbuzzShouldReturnFizzBuzzForFifteen(t *testing.T) {
	result := fizzbuzz(15)

	assert.Equal(t, "fizzbuzz", result)	
}

func TestFizzBuzzShouldReturnFizzBuzzForThirty(t *testing.T) {
	result := fizzbuzz(30)

	assert.Equal(t, "fizzbuzz", result)
}

func TestFizzBuzzShouldReturnNumberForZero(t *testing.T) {
	result := fizzbuzz(0)

	assert.Equal(t, "0", result)
}
func TestFizzbuzzShouldReturnFizzForNumbersDivisibleByNegativeThree(t *testing.T) {
	result := fizzbuzz(-3)

	assert.Equal(t, "fizz", result)
}