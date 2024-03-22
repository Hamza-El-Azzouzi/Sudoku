<h1>Sudoku Solver in Go</h1>

<p>This is a Sudoku solver implemented in Go. It uses a backtracking algorithm to efficiently solve Sudoku puzzles of varying difficulties.</p>

<h2>Usage</h2>

<p>To run the Sudoku solver, execute the following command in your terminal:</p>

<pre><code>go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e</code></pre>

<p>The provided arguments represent the Sudoku puzzle rows. Each row is passed as a separate argument in the form of a string where empty cells are denoted by dots (<code>.</code>).</p>

<h2>Example</h2>

<h3>Input:</h3>

<pre><code>.96.4...1
1...6...4
5.481.39.
..795..43
.3..8....
4.5.23.18
.1.63..59
.59.7.83.
..359...7</code></pre>

<h3>Output:</h3>

<pre><code>
3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$
</code></pre>

<h2>Code Structure</h2>

<p>The code is structured as follows:</p>

<ul>
<li><strong>Constants and Variables</strong>: Definitions of constants like <code>gridSize</code> and functions to print the grid, find empty locations, and check for the validity of Sudoku puzzles.</li>
<li><strong>Sudoku Solving Algorithms</strong>: Functions for solving Sudoku puzzles, including backtracking and helper functions to check if a number is safe to place in a given position.</li>
<li><strong>Main Function</strong>: Reads the Sudoku puzzle from command-line arguments, solves it, and prints the solution.</li>
</ul>

<h2>How It Works</h2>

<p>The solver utilizes a backtracking algorithm to fill in the Sudoku grid recursively. It tries different numbers in empty cells and backtracks if a solution cannot be found.</p>

<h2>Contributing</h2>

<p>Contributions are welcome! If you have any suggestions, bug reports, or enhancements, feel free to open an issue or submit a pull request.</p>

<h2>License</h2>

<p>This project is licensed under the MIT License - see the <a href="LICENSE">LICENSE</a> file for details.</p>
