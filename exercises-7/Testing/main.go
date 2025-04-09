package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
    result := add(2, 3)
    assert.Equal(t, 5, result, "The result should be 5")
}


func add(a int, b int) int {
    return a + b 
}


func TestDivision(t *testing.T) {
    _, err := divide(5, 0)
    assert.Error(t, err, "An error should occur when dividing by zero")
}


func divide(a int, b int) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return float64(a) / float64(b), nil
}

func TestSubtraction(t *testing.T) {
	result := subtract(5, 3)
	assert.Equal(t, 2, result, "The result should be 2")
}	

func subtract(a int, b int) int {
	return a - b
}

func TestMultiplication(t *testing.T) {
	result := multiply(2, 3)
	assert.Equal(t, 6, result, "The result should be 6")
}

func multiply(a int, b int) int {
	return a * b
}	

func main() {
	fmt.Println("Running tests...")

	testing.Main(func(pat, str string) (bool, error) { return true, nil }, []testing.InternalTest{
		{"TestAddition", TestAddition},
		{"TestDivision", TestDivision},
		{"TestSubtraction", TestSubtraction},
		{"TestMultiplication", TestMultiplication},
	}, nil, nil)


}