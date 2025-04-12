# mwpm
This is a fork of [HyosangKang/mwpm](https://github.com/HyosangKang/mwpm) with a little modification to profile performance.
It solves the minimum weight perfect matching problem using the Blossom algorithm.
some SolverI are supported:

1. MaxSolver to solve maximum weight perfect matching problem
2. FilterSolver to cut less promising edges to improve performance. it might lead to 
   suboptimal solution or even no solution.
