### Day 8: Binary search

My today’s algorithm is [binary search](https://en.wikipedia.org/wiki/Binary_search_algorithm) on a sorted array. This one is fascinating, since it is probably one of the easiest algorithms to understand, yet one of the most difficult to implement correctly.

According to Sedgewick, it was invented in mid-50s, but it was not until mid-60s for the first correct implementation to appear. In 2006 binary search routine in Java libraries had to be fixed due to bug in its implementation.

Why is that? You have to be very careful when playing with indices. The key to success is to keep `(left, right)` boundaries within the searched range rather than outside.

#### Algorithm

Given an array A of n elements with values or records `A0, A1, ..., An−1`, sorted such that `A0 ≤ A1 ≤ ... ≤ An−1`, and target value `T`, the following subroutine uses binary search to find the index of `T` in `A.[7]`

- Set L to 0 and R to n − 1.
- If L > R, the search terminates as unsuccessful.
- Set m (the position of the middle element) to the floor (the largest previous integer) of `(L + R) / 2`.
- If Am < T, set L to m + 1 and go to step 2.
- If Am > T, set R to m − 1 and go to step 2.
- Now Am = T, the search is done; return m.

This iterative procedure keeps track of the search boundaries with two variables. Some implementations may check whether the middle element is equal to the target at the end of the procedure. This results in a faster comparison loop, but requires one more iteration on average.
