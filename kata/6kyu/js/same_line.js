function onLine(points) {
    // formula for coline
    // y2 - y1 / x2 - x1 === y3 - y2 / x3 - x2

    let dif = 0;

    for (let i = 1; i < points.length; i++) {
        let x = (points[i][1] - points[i - 1][1]) / (points[i][0] - points[i - 1][0]);
        if (dif === 0) {
            dif = x;
            continue;
        }

        if (x !== dif) {
            return false;
        }
    }

    return true;
}

console.log(onLine([[1, 2], [7, 4], [22, 9]])) // true
console.log(onLine([[1, 2], [-3, -14], [22, 9]])) // false
console.log(onLine([[1, 4], [1, 4], [3, 4], [17, 4], [34, 4]])) // true
console.log(onLine([[9, -3], [9, 3], [9, 29], [9, 11], [9, 25], [9, 13], [9, -5]])) // true
console.log(onLine([[-5, 23], [-5, -77], [-5, 73], [-5, 53], [-5, -7], [-5, 43]])) // true
