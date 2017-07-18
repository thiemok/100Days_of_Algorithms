### Day 6: Postfix notation
While humans mostly use infix notation of algebraic expressions, Reverse Polish notation or postfix notation is much easier to parse algorithmically.
For example, these expressions are equivalent, the first one is in infix form and the second one is in postfix form.
`2 * (1 + 3) = 2 1 3 + *`

#### Algorithm

- While there are input tokens left
	- Read the next token from input.
	- If the token is a value
	- Push it onto the stack.
	- Otherwise, the token is an operator (operator here includes both operators and functions).
	- It is already known that the operator takes n arguments.
	- If there are fewer than n values on the stack
		- (Error) The user has not input sufficient values in the expression.
	- Else, Pop the top n values from the stack.
		- Evaluate the operator, with the values as arguments.
		- Push the returned results, if any, back onto the stack.
- If there is only one value in the stack
	- That value is the result of the calculation.
- Otherwise, there are more values in the stack
	- (Error) The user input has too many values.
