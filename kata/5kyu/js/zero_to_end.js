function moveZeros(arr) {
    let nonZeros = [];
    let zeros = [];

    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === 0) {
            zeros.push(arr[i]);
            continue;
        }
        nonZeros.push(arr[i]);
    }
    return nonZeros.concat(zeros);
}

console.log(moveZeros([
    9, 0, 0, 9, 1, 2, 0,
    1, 0, 1, 0, 3, 0, 1,
    9, 0, 0, 0, 0, 9
]));
