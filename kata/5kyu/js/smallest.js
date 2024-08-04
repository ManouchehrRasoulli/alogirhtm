function smallest(n) {
    let row = String(n).split("");
    let min = [n, 0, 0],
        test = [];

    let length = row.length;
    for (let i = 0; i < length; i++) {
        for (let j = 0; j < length; j++) {
            test = row.slice(0, i).concat(row.slice(i + 1));

            if (Number(test.slice(0, j).concat(row[i], test.slice(j)).join("")) < min[0]) {
                min = [Number(test.slice(0, j).concat(row[i], test.slice(j)).join("")), i, j];
            }
        }
    }

    return min;
}

console.log(smallest(285365));