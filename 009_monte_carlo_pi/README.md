### Day 9: Monte carlo - π
I guess I’ll introduce more the one randomised algorithm in this series, and that’s not only because I love probabilistic approach.
Randomised simulation is often the best/only way to solve otherwise intractable problems.

_if you are mathematician, excuse my sloppy wording; randomised simulation should be understood as a random sampling from space of simulations under given distribution_

How can we estimate π if the only tool we have at disposal is a good random number generator? When we choose a random coordinate (x, y) in range (-1, 1) and each point has equal chance to be chosen, the probability to hit a circle with unit radius is

![Monte carlo pi](https://cdn-images-1.medium.com/max/1600/1*uY6ljvUZpR9wNEgUPJfEFw.png)

Having sufficiently large set of points [and a good generator] we can get as close as we want according to [Chebyshev’s inequality](https://en.wikipedia.org/wiki/Chebyshev%27s_inequality)

#### Algorithm

-  For a number of samples N, do N times
	- Pick a random Point `(x, y)` within `[-1, 1]`
	- If the point is within the unit circle, increase count M
- retrun `4 * M / N`
