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
            return loc;
        case down:
            loc.x++;
            return loc;
        case left:
            loc.y--;
            return loc;
        case up:
            loc.x--;
            return loc;
    }
}

function createSpiral(n) {
    if (typeof n !== "number" || n < 1) {
        throw [];
    }
    if (n === 1) {
        return [1];
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

    for (let i = 1; i <= n + 4; i++) {
        console.log(loc, dir);
        matrix[loc.x][loc.y] = i;
        nextPos(loc, dir);
        console.log("nx", loc, dir);
        if (loc.x === n || loc.y === n || loc.x === -1 || loc.y === -1) {
            switch (dir){
                case right:
                    loc.x--;
                    break;
                case down:
                    loc.y--;
                    break;
                case left:
                    loc.x++;
                    break;
                case up:
                    loc.y++;
                    break;
            }
            dir = changeDirection(dir);
            nextPos(loc, dir);
            console.log("change direction");
        }
    }

    return matrix;
}

console.log(createSpiral(3));
