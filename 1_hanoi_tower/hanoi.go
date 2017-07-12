package hanoi

import "github.com/emirpasic/gods/stacks/arraystack"

/**
 * Solves the towers of hanoi algorithm for n Disks.
 * Returns final result and number of steps for testing purposes
 * @param		{uint}			n					Number of disks
 * @return	{[3]*Stack}	result		Final result of the algorithm
 * @return	{uint}			numSteps	Number of steps it took to solve the algorithm
 */
func Hanoi(n uint) (result [3]*arraystack.Stack, numSteps uint) {
	var stacks = initStacks(n)

	numSteps = hanoi(n, stacks[0], stacks[2], stacks[1])

	return stacks, numSteps
}

func initStacks(n uint) (stacks [3]*arraystack.Stack) {
	for stack := range stacks {
		stacks[stack] = arraystack.New()
	}
	var i uint
	for ; i < n; i++ {
		stacks[0].Push(n - i)
	}
	return stacks
}

func movePeg(source *arraystack.Stack, target *arraystack.Stack) {
	peg, _ := source.Pop()
	target.Push(peg)
}

func hanoi(nDisks uint, source *arraystack.Stack, target *arraystack.Stack, spare *arraystack.Stack) (numSteps uint) {
	/**
	 * Step 4: no disks left to move
	**/
	if nDisks == 0 {
		return
	}
	/**
	 * Step 1: Move N − 1 disks from the source to the spare peg,
	 * by the same general solving procedure. Rules are not violated, by assumption.
	 * This leaves the disk N as a top disk on the source peg.
	**/
	numSteps += hanoi(nDisks-1, source, spare, target)

	/**
	 * Step 2: Move the disk N from the source to the target peg,
	 * which is guaranteed to be a valid move,
	 * by the assumptions — a simple step.
	**/
	movePeg(source, target)
	numSteps++

	/**
	 * Step 3: Move the N − 1 disks that we've just placed on the spare,
	 * from the spare to the target peg by the same general solving procedure,
	 * so they are placed on top of the disk N without violating the rules.
	 */
	numSteps += hanoi(nDisks-1, spare, target, source)

	return numSteps
}
