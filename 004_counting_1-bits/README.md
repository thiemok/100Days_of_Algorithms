### Day 4: Counting 1-bits
When you write down a positive integer as binary number, how many 1s do you get? For example, `99 = 0b1100011`
An obvious solution would be to take a single-bit mask, walk over the number `X & (1 << k)` and count non-zero results up. But this solution counts both 0s and 1s.

Can you do better by only counting 1s and completely ignoring 0s? There’s a ton of beautiful and surprising bit-tricks you can do. Too bad we don’t need them any more.
This is also known as the [Hamming weight](https://en.wikipedia.org/wiki/Hamming_weight) or popcount.

#### Algorithm
This is going to be a relativly simple and readable aproach.
In any case where performance could be a factor one should use a built-in implementation like gcc's `int __builtin_popcount (unsigned int x)` instead.

1. Check if lowest bit is 1 and if so increase count
2. Shift input 1 bit right
3. Repeat until input is 0
