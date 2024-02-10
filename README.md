# z01-tetris-optimizer

Tetris-optimizer is an optimization problem. Given several tetris pieces, you have to find a piece arrangement that fits in the smallest square possible. Detailed subject can be found in subject.md.

## How to use

`go run . "file"` to use algorithm.
`./test.sh` to run tests on provided examples. b* files need to return error, while other files should return a solution.

## Project state

The implemented algorithm uses recursive bruteforce to find a solution. Details can be found in code comment. The final solution is colored using DSatur algorithm for easier reading.

Performance wise, the hardest given example is solved in less than 2 seconds, which looks good, especially compared with some other student's solutions.