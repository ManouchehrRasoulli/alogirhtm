function _neighbors(matrix, n, point) {
    let ng = [[point[0] - 1, point[1]],
        [point[0], point[1] - 1],
        [point[0] + 1, point[1]],
        [point[0], point[1] + 1]];

    let final = [];
    for (let i = 0; i < 4; i++) {
        if (ng[i][0] < 0 || ng[i][0] > n || ng[i][1] < 0 || ng[i][1] > n) {
            continue;
        }
        if (matrix[ng[i][0]][ng[i][1]] === 'W') {
            continue;
        }
        final.push(ng[i]);
    }

    return final;
}

function pathFinder(maze) {
    // let mazeMatrix = toMatrix(maze);
    let mazeMatrix = maze.split(`\n`).map(l => l.split(``));
    let n = mazeMatrix.length - 1;

    let q = [[0, 0]];
    mazeMatrix[0][0] = "0"

    while (q.length !== 0) {
        let current = q.shift();
        let currentCost = Number(mazeMatrix[current[0]][current[1]]);
        if (current[0] === n && current[1] === n) {
            return Number(mazeMatrix[n][n]);
        }
        let ngs = _neighbors(mazeMatrix, n, current);
        for (let i = 0; i < ngs.length; i++) {
            let v = Number(mazeMatrix[ngs[i][0]][ngs[i][1]]);
            if (isNaN(v)) {
                mazeMatrix[ngs[i][0]][ngs[i][1]] = String(currentCost + 1);
                q.push(ngs[i]);
                continue;
            }
            if (v+1 < currentCost) {
                mazeMatrix[current[0]][current[1]] = String(v);
                q.push(ngs[i]);
            }
        }
    }

    return false;
}


let maze = `...W...W
W......W
........
WWW.....
..W..W..
WW.W....
W.WWW..W
...W....`;

console.log(pathFinder(maze));