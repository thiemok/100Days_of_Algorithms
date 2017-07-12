### Day 1: Hanoi Tower

The Tower of Hanoi (also called the Tower of Brahma or Lucas' Tower and sometimes pluralized) is a mathematical game or puzzle. It consists of three rods and a number of disks of different sizes, which can slide onto any rod. The puzzle starts with the disks in a neat stack in ascending order of size on one rod, the smallest at the top, thus making a conical shape.

The objective of the puzzle is to move the entire stack to another rod, obeying the following simple rules:
- Only one disk can be moved at a time.
- Each move consists of taking the upper disk from one of the stacks and placing it on top of another stack.
- No disk may be placed on top of a smaller disk.

With 3 disks, the puzzle can be solved in 7 moves. The minimal number of moves required to solve a Tower of Hanoi puzzle is 2^n − 1, where n is the number of disks.

#### Algorithm
To move N disks from left to right:
1. Move N − 1 disks from the source to the spare peg, by the same general solving procedure. Rules are not violated, by assumption. This leaves the disk N as a top disk on the source peg.
2. Move the disk N from the source to the target peg, which is guaranteed to be a valid move, by the assumptions — a simple step.
3. Move the N − 1 disks that we've just placed on the spare, from the spare to the target peg by the same general solving procedure, so they are placed on top of the disk N without violating the rules.
4. The base case being to move 0 disks, that is, do nothing – which obviously doesn't violate the rules.
