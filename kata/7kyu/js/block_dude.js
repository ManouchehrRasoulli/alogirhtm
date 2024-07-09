function canTraverse(matrix) {
    let d = [0];
    for (let i = 0; i < 9; i++) {
        let boxCount = 0;
        for (let j = 3; j >= 0; j--) {
            if (matrix[j][i] !== 1) {
                break;
            }
            boxCount++;
        }
        d[i+1] = boxCount;
    }

    console.log(d);

    for (let i = 0; i < 9; i++) {
        let dif = d[i] - d[i+1];
        if (dif < -1 || dif > 1) {
            return false;
        }
    }

    return true;
}

console.log(canTraverse([
    [0, 0, 0, 0, 0, 0, 0, 1, 0],
    [0, 0, 0, 1, 1, 0, 1, 1, 1],
    [1, 0, 1, 1, 1, 1, 1, 1, 1],
    [1, 1, 1, 1, 1, 1, 1, 1, 1]
]));