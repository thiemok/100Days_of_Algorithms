### Day 2: Matrix chain multiplication
Matrix multiplication is an associative operation. `(AB)C` is equal to `A(BC)` for matrices `A, B, C`, and it doesn’t matter which order of pair multiplications you choose.

Unfortunately, that’s not true from computational perspective.
If dimensions of matrices are `A=[10, 20]`, `B=[20, 30]`, `C=[30, 40]`, numbers of scalar multiplications differ significantly:
```
(AB)C = 10*20*30 + 10*30*40 = 18000
A(BC) = 20*30*40 + 10*20*40 = 32000
```
The best ordering can be found using recursive relationship. Let MCM denote a function that returns a minimum number of scalar multiplications. Then MCM can be defined as the best split among all possible choices.

<img class="graf-image" data-image-id="1*ZvVBe6tFLKmds_oDe-nzyw.png" data-width="802" data-height="42" data-action="zoom" data-action-value="1*ZvVBe6tFLKmds_oDe-nzyw.png" src="https://cdn-images-1.medium.com/max/1600/1*ZvVBe6tFLKmds_oDe-nzyw.png" style="">

Using dynamic programming and memoization, the problem can be solved in O(n³) time.

#### Algorithm

1. Take the sequence of matrices and separate it into two subsequences.
2. Find the minimum cost of multiplying out each subsequence.
3. Add these costs together, and add in the cost of multiplying the two result matrices.
4. Do this for each possible position at which the sequence of matrices can be split, and take the minimum over all of them.
