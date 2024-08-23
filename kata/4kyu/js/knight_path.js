const chestBoard = [
    ["a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8"],
    ["a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7"],
    ["a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6"],
    ["a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5"],
    ["a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4"],
    ["a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3"],
    ["a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2"],
    ["a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1"]
]

function getRowCol(pos) {
    for (let i = 0; i < chestBoard.length; i++) {
        let j = chestBoard[i].indexOf(pos);
        if (j !== -1) {
            return [i, j];
        }
    }
    return undefined;
}

function getPos(row, col) {
    if (row < 0 || col < 0 || row > 7 || col > 7) {
        return undefined;
    }

    return chestBoard[row][col];
}

function nextMove(pos) {
    let current = getRowCol(pos);
    if (current === undefined) {
        return undefined;
    }

    let row = current[0];
    let col = current[1];

    let next = [[row - 2, col + 1], [row - 2, col - 1], [row + 2, col + 1], [row + 2, col - 1], [row - 1, col + 2], [row - 1, col - 2], [row + 1, col + 2], [row + 1, col - 2],]

    let moves = [];
    for (let i = 0; i < next.length; i++) {
        let pos = getPos(next[i][0], next[i][1]);
        if (pos === undefined) {
            continue;
        }

        moves.push(pos);
    }

    return moves;
}

function knight(start, finish) {
    let q = [];
    let path = [];
    q.push([start, 0]);
    let min = [finish, 1000];

    while (q.length !== 0) {
        let current = q.shift();
        if (path.indexOf(current[0]) !== -1) {
            continue;
        }
        path.push(current[0]);

        if (current[0] === finish && current[1] < min[1]) {
            min[1] = current[1];
        }

        let moveCost = current[1] + 1;
        if (moveCost > min[1]) {
            continue;
        }

        let moves = nextMove(current[0]);
        for (let i = 0; i < moves.length; i++) {
            q.push([moves[i], moveCost]);
        }
    }

    return min[1];
}

console.log(knight("a1", "c1")); // 2
console.log(knight("a1", "f1")); // 3
console.log(knight("a1", "f3")); // 3
console.log(knight("a1", "f4")); // 4
console.log(knight("a1", "f7")); // 5
