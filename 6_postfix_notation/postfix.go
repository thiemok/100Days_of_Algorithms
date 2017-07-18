package postfix

import (
	"errors"
	"math"
	"strconv"
	"strings"

	aStack "github.com/emirpasic/gods/stacks/arraystack"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/**
 * Solves the given term in postfix notation
 * @param 	{string}	input	The term to solve, needs to be in postfix notation and delimited by spaces
 * @return	{int}			The result
 * @return	{error}		Error if the input was invalid
 */
func Solve(input string) (int, error) {

	operations := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
		"^": pow,
	}

	stack := aStack.New()

	for _, token := range strings.Split(input, " ") {
		value, error := strconv.Atoi(token)
		if error == nil {
			stack.Push(value)
		} else {
			op, ok := operations[token]
			if !ok || stack.Size() < 2 {
				return 0, errors.New("Invalid input")
			}
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(op(a.(int), b.(int)))
		}
	}
	if stack.Size() != 1 {
		return 0, errors.New("Invalid input")
	}
	res, _ := stack.Pop()
	return res.(int), nil
}
