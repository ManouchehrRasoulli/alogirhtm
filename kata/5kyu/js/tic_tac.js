function isSolved(board) {

    function checkHor() {
        var res;
        board.forEach(function (item) {
            if (item[0] == item[1] &&
                item[0] == item[2] &&
                item[1] == item[2] &&
                item[0] != 0)
                res = item[0];
        });
        return res;
    }

    function checkVer() {
        for (var a = 0; a < 3; a++) {
            if (board[0][a] == board[1][a] &&
                board[1][a] == board[2][a] &&
                board[2][a] == board[0][a] &&
                board[0][a] != 0)
                return board[0][a];
        }
    }

    function checkA1() {
        if ((board[0][0] == board[1][1] &&
                board[1][1] == board[2][2] &&
                board[2][2] == board[0][0]) ||
            (board[0][2] == board[1][1] &&
                board[1][1] == board[2][0] &&
                board[2][0] == board[0][2]))
            return board[1][1];
    }


    function gameOver() {
        for (var x = 0; x < 3; x++)
            for (var y = 0; y < 3; y++)
                if (board[x][y] == 0)
                    return -1;
        return 0;
    }

    return checkHor() || checkVer() || checkA1() || gameOver();
}

console.log(isSolved([
    [0, 0, 1],
    [0, 1, 2],
    [2, 1, 0]]) === -1)