/**
 * @param {number[][]} grid
 * @param {number} i
 * @param {number} j
 */
var minVal = function (grid, i, j) {
    let m = grid.length;
    let n = grid[m - 1].length;
    if (i + 1 >= m && j + 1 >= n) {
        return;
    }

    if (i + 1 >= m) {
        grid[i][j] += grid[i][j + 1];
        return;
    }

    if (j + 1 >= n) {
        grid[i][j] += grid[i + 1][j];
        return;
    }

    grid[i][j] += Math.min(grid[i + 1][j], grid[i][j + 1]);
    return;
};

/**
 * @param {number[][]} grid
 * @return {number}
 */
var minPathSum = function (grid) {
    let m = grid.length;
    let n = grid[m - 1].length;

    for (let i = m - 1; i >= 0; i--) {
        for (let j = n - 1; j >= 0; j--) {
            minVal(grid, i, j);
        }
    }

    console.log(grid);

    return grid[0][0];
};


console.log(minPathSum([[1, 3, 1], [1, 5, 1], [4, 2, 1]])); // 7
console.log(minPathSum([[1, 2, 3], [4, 5, 6]])); // 12
