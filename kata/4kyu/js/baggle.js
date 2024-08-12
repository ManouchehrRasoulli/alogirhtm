Array.prototype.containsPoint = function (point) {
    for (let i = 0; i < this.length; i++) {
        if (this[i][0] === point[0] && this[i][1] === point[1]) {
            return true;
        }
    }
    return false;
};

function _neighbors_(n, point) {
    let neighbors = [
        [point[0] - 1, point[1]], [point[0], point[1] - 1],
        [point[0] + 1, point[1]], [point[0], point[1] + 1],
        [point[0] - 1, point[1] - 1], [point[0] - 1, point[1] + 1],
        [point[0] + 1, point[1] + 1], [point[0] + 1, point[1] - 1]];

    let final = [];
    for (let i = 0; i < 8; i++) {
        if (neighbors[i][0] < 0 || neighbors[i][0] >= n || neighbors[i][1] < 0 || neighbors[i][1] >= n) {
            continue;
        }
        final.push(neighbors[i]);
    }

    return final;
}

function findWord(board, point, word, index, path) {
    if (path.containsPoint(point)) {
        return false;
    }

    if (board[point[0]][point[1]] !== word[index]) {
        return false;
    }

    if (index === word.length - 1) {
        return true;
    }

    path.push(point);

    let ng = _neighbors_(board.length, point);
    for (let i = 0; i < ng.length; i++) {
        if (findWord(board, ng[i], word, index + 1, path)) {
            return true;
        }
    }

    return false;
}

function checkWord(board, word) {
    board = board.map(x => x.slice());
    for (let i = 0; i < board.length; i++) {
        for (let j = 0; j < board[i].length; j++) {
            if (board[i][j] !== word[0]) {
                continue;
            }
            if (findWord(board, [i, j], word, 0, [])) {
                return true;
            }
        }

    }

    return false;
}

/*
let testBoard = [
    ["E", "A", "R", "A"],
    ["N", "L", "E", "C"],
    ["I", "A", "I", "S"],
    ["B", "Y", "O", "R"]
];
 */

let testBoard = [
    ["T", "T", "M", "D", "A"],
    ["G", "Y", "I", "N", "N"],
    ["P", "A", "L", "C", "E"],
    ["I", "A", "U", "L", "G"],
    ["A", "M", "I", "N", "A"]
];

console.log(checkWord(testBoard, "ANIMA"));
