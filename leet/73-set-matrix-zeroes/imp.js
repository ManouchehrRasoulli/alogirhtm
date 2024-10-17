/**
 * @param {number[][]} matrix
 * @param {number} row
 * @param {number} col
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
var setNegate = function (matrix, row, col) {
    for (let i = 0; i < matrix.length; i++) {
        if (matrix[i][col] === "-1") {
            break;
        }

        if (matrix[i][col] !== 0) {
            matrix[i][col] = "-1";
        }
    }
    for (let j = 0; j < matrix[0].length; j++) {
        if (matrix[row][j] === "-1") {
            break;
        }

        if (matrix[row][j] !== 0) {
            matrix[row][j] = "-1";
        }
    }
}

/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
var setZeroes = function (matrix) {
    for (let i = 0; i < matrix.length; i++) {
        for (let j = 0; j < matrix[i].length; j++) {
            if (matrix[i][j] === 0) {
                setNegate(matrix, i, j);
            }
        }
    }

    for (let i = 0; i < matrix.length; i++) {
        for (let j = 0; j < matrix[i].length; j++) {
            if (matrix[i][j] === "-1") {
                matrix[i][j] = 0;
            }
        }
    }
};


let mx = [[1, 1, 1], [1, 0, 1], [1, 1, 1]];
setZeroes(mx);
console.log(mx);

mx = [[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]];
setZeroes(mx);
console.log(mx);

mx = [[0, 1]];
setZeroes(mx);
console.log(mx);
