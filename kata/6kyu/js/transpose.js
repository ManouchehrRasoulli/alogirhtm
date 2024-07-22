function transpose(matrix) {
    let rows = matrix.length;
    let cols = matrix[0].length;

    let result = [];
    for(let j = 0; j < cols; j++) {
        result[j] = [];
    }

    for(let i = 0; i < rows; i++) {
        for(let j = 0; j < cols; j++) {
            result[j][i] = matrix[i][j];
        }
    }

    return result;
}

console.log(transpose([[1]]));
console.log(transpose([[1, 2, 3], [4, 5, 6]]));
