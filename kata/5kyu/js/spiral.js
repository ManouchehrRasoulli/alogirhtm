const right = "r";
const down = "d";
const left = "l";
const up = "u";

function changeDirection(dir) {
    switch (dir) {
        case right:
            return down;
        case down:
            return left;
        case left:
            return up;
        case up:
            return right;
    }
}

function nextPos(loc, dir) {
    switch (dir) {
        case right:
            loc.y++
            break;
        case down:
            loc.x++;
            break;
        case left:
            loc.y--;
            break;
        case up:
            loc.x--;
            break;
    }
}

function createSpiral(n) {
    if (Math.round(n) !== n || n < 1) {
        return [];
    }
    if (n === 1) {
        return [[1]];
    }

    let matrix = [];
    for (let i = 0; i < n; i++) {
        matrix[i] = new Array(n);
    }

    let loc = {
        x: 0,
        y: 0,
    };

    let dir = right;

    for (let i = 1; i <= n * n; i++) {
        console.log(loc, dir);
        matrix[loc.x][loc.y] = i;
        let nextLoc = {
            x: loc.x,
            y: loc.y
        };
        nextPos(nextLoc, dir);
        console.log("nx", nextLoc, dir);
        if (nextLoc.x > n - 1 || nextLoc.y > n - 1 || nextLoc.x < 0 || nextLoc.y < 0) {
            dir = changeDirection(dir);
            nextPos(loc, dir);
            continue;
        }
        if (matrix[nextLoc.x][nextLoc.y] !== undefined) {
            dir = changeDirection(dir);
            nextPos(loc, dir);
            continue;
        }
        loc = nextLoc;
    }

    return matrix;
}

console.log(createSpiral(3));
console.log(createSpiral(4.5));
