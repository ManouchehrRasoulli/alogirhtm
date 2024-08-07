/*
Array.prototype.containsPoint = function (point) {
    for (let i = 0; i < this.length; i++) {
        if (this[i][0] === point[0] && this[i][1] === point[1]) {
            return true;
        }
    }

    return false;
}
 */

function toMatrix(maze) {
    console.log(maze);
    let mz = [];
    let rows = maze.split('\n');
    for (let i = 0; i < rows.length; i++) {
        mz.push(rows[i].split(''));
    }
    console.log(mz);
    console.log(mz.length);
    return mz;
}

function neighbors(matrix, n, point) {
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

    while (q.length !== 0) {
        let current = q.pop();
        mazeMatrix[current[0]][current[1]] = 'W';
        if (current[0] === n && current[1] === n) {
            return true;
        }
        q.push(...neighbors(mazeMatrix, n, current));
    }

    return false;
}


let m = `..W..WW.W.........W............W......W..WW.WW.WW....W..........WW...............W.W.....W..WW......
...W.....WW.........W...W........WW......W.W.W.W.W.W.WW....W..W....W..WWW..........W..W..........W..
.WW.......W.WW.WWW....W...W..........W.W.....W.....W..............W.....W......W...W..W....W........
....W..W...WW..WW.............W....W.....WW..W.....WW......WW.......W.WWWW.....W..WW.....WW.........
.W.W..W...WW.W..W.WWWW.W.........WW....W.W......WW...........W.....W....W...W..W.............WW....W
W...WW...W.W.W.....W.WW...W..WW...W....WW...........W..W.WW.W....WW...W.W..WW....W.W....W..W......WW
.WW..........WW.W.W....W..WW...........W..............W...WWWW.........WWW.W....WW.W....W.W....W....
..WWW...W..WWW...W.W.WWW..........W..W..............W...W.............WW.........W......W.WWW..W....
...WWW.W....W...W....W..WWWW....WW.W...........WW..W........W...W....W.......WW.W....W.....W.....W.W
..WWWWW.....W.....WW.W.WW..WWW.....W.......W...WW.W...W..W...W..W.W...W.....W..W.W........WWW......W
.WW.......W......W....W..W.WW.......W.WWW.WW....W.W.W....WW..WW.W.WW..W...W.W.....W...WW.W..W....W.W
W.WW....W..W..............W.......WWWWW...W.WW.....WW..W......W.W.W..........WWW...W..W.W...........
.......WW..WWW....W.............W.W...WW..W............W...WW.WW.W........W...W..W..W.W..W...W......
.W......WW.W..WWW..........W...W.W...W.W.W.W..............W..W.............WW...W..WW...W.....WW....
.W..W..WW..W.........W....W....WW.W.W..W...W.......WW.W.......W......W..WWWW.....W..WW.WW...W.......
..W...W..W........W..W....W..........W.....W...W.W..WWW..WW.W.W..W..WWW....W.WW...W.....W..W.W.W....
......W.....WW....W....WWW..WWW..W.....WWWWWW.W.....WW.WW.........WW......WW............WW...WWWWW.W
....W..WW.....WWW.W...WW....W.......W.....W.WW.WW...WW.W........WWW..W....WWW.WW.......W.W...WW.....
W.W.....WWW..W.W.WWWWW.WWW.W..W...W.......WWW..WW.W.....W..W.W..WW...W.......W........WW..W.W.WW..W.
W.W..WW....W.W.WWWWWW....W.....W..W.WW....WW...W.WW..W.......W.....W....W.W....W.W..W.WW.W...W....W.
........WW.W..W..W..WWW.........W...WW...W....W..W...W.WW...W............W....WW.....W.W....WW....WW
...W....W...W.......W..W..W...........W.WWWWW..W..WW........W..WWW....WWW...W....W.......W.......W..
...W.....WW.......W....W..W.WW.W....W....W..WWWW.W.W.W.WW..W....W.WWWW..........W...W..WW..W.....W..
W.....W...WW.W........W.WW...W....W.W.......W...W.W..W....WW.....W.....W.W...W...........W...WW.W...
W.......W.W.........W.W........W...W...W......W......W..WW....W.W...W.W....W...............W.W......
.W.W..W.W.W...........WW..W...W......W..W.W..........WW.WWW.....WW.......W.W..W....W.W.....W...W....
W........WWW...W.W.W......W...W...WW.....W.W.....WWW.....W.W.......W.......WW...W.W......W...WW.....
W.......WW.W.WW..W....W.W...WWW....WW...W......W.W..W.W..W.WW...W.....WW.WWW..W.W.....W...W....WW..W
..........WW..W.........W.....W.W........W..W.W..W.......W...WW..W...........W.WW.WW....W......WW.W.
.WWW.......W.........W...............W......W.W....W.W..W...W.W....W....W.W.W..........W.WW.....W...
.W..W..W..W.WWW............W..WW..........W......W...W......W....W.W.........W...W...W.W....W.WWWWWW
W...WWW.W..W...W.W........W.....W..W.W.WW.W.....W.W..WW..W...W.WW.W.....W.W..W...W...W...W......W...
WW..WWW...W...........W.....WW....W..W...W.W.....W......W.....WW......W..W...WW.W..WW....WW.W.W.....
W.WW.W...W...WWW..W.W.....WWW.W...W...W...W......W....WWWWW.W........WWW..WW..W...W.W.....W.W.......
..WW....WW.WW.....W.W..W.W...........WW......W.WW.....W....W...W..........W..W.....W...WWW.W.W..W...
.W.W..W.....W..W.........WW......WW..W...........W..WW.W..W.............WWW..W.W.......WW....WW.W...
.W......W..WW..W.W..W..............W..W.WWW.W...WW...WW.W.WW.WW...W.W..W.W....W......W......W..W...W
W........W...W.W...W..W.W...WW.WW........W..W........W.W...W.......W..WW..WW....W....W....W.....WW..
........W.W...W.W....WW.............W.WW..W..WW.....W.....W.W...WW..W.......W....W.....W..W.W.WW..W.
........W..WWWW.W..W...W...W.W.....W.....W..WW..W..W......................WW.................W....W.
.W.W.WWWWW..WW.W......W.W....W....W....W.......W.....W.W........WWW..........W.W.W..W....WW.W.W..W..
..W.W.WWW..W....W..WW...WW.......W.W.W...........W..W...W..WWW..W...WWW........W..W.W..W.W......W..W
..WW...W...........W...W..W.W...W........WW....WW..WW..W.WW.W.....W...W..WW...WW..W.....W..W.WW.W.W.
..W.W...........W........W..W.W....WW...W.W..WW.......W.....W....WWW.W.........W.....W..W...WW...WW.
......WWW...W.W..W...W..W....W.....W.WW.....W.WW....WWW....W.W..W.......WW.......W.....W.......WW...
.....W.WW..W...W.......W.....W.............W...W.WW..WW....WW....WW...WW.......WWW..W.........W..W.W
W...W........W.....WWWWW..W..WW.W.W..W...W..W...W...W....W.W..........W.........W......W.W....W.W..W
W..WW..W..........W..........W...W.WW......W....W....WW...WWW..WW.....W.WW..W.W...W........WW...W...
...W.WW..WW.W........W....WWW..WW...W.WW......WW............W....W.....WW...W..W...W..W...W........W
....W.W.W.....WWW........WW.........WW...W.WW.....W....W...WW...W.....W..W....W.W..W.WW.WWW..W..WW..
WWW.W..WW.WW..WW.W.WW...WWW...W..WWW.....WW.WW..W....W.W.....WW.W......W.WWW..W...WW......W..W...W.W
....W.......W...W..W...W.W.W.W.....W.....W.W...W.W.W..W.W.......W................W.....W.W..W.W.W...
W..W...W.W.W....W....W.W.W..W.......WWW..W...W..WW..W.WW.W.W......W.W..W...W..W.W...W.W.WW..W.W....W
..............W.W...W.W.W..WW.....W....W..W....W..W..W.........W..W.W.W...W.W.W.W...W..W.W.W......W.
W.W.W.....W.WW.W..W...W....WWW.W.....W...W....WWW.W.....W.W.....W..W...W..W.W...WW..W.........W.....
W...W.W....W.W..W.W.....W..W..W...W...WW...WW....W..WW.W.W.......W.W.W...W....W...W.W.W......W......
WWWW..W..W.......W.WW...WW...W......WW...W.W.....W...W.....W...W.W.W......W.WWW..W....WW............
..W........W.WWW..W.W.WWWWW..W.WW.W........W..WW.......W..W......W...W.W..WW.W........W....W.....W.W
..WW..W.W.W...WW.......W..W....W..WWW......................WW...WW.....W......W....W....W.W..W......
......W.WWWW..W.....W........WWW...........WW..W.W..WW.W...WW.W...W......W....W.W.W....WWWWWW.W....W
.W.W...W..W.WW.W.W......W..WWWW.W.W.W...W.........W...WW..W.W.W...WW..W.WWW..W.W...........W.WW...WW
WWW..........W.WWW.....W....W....WW..W.......W.W.W..WW...W.WW.W...W..W..........W..W.W.W......W....W
.....W.....W......WW...WW.......W.......WW.......W....W...W.WW..W.W.WW.W.WW....W.........W......W...
W...WW....WWWW...WW...W....WW.WWWW........W..WW.W...WW..WWW..............W....W..W.....W.WW...W.W.W.
W.WW.W.W..WW.W.W.........W...W......W....W.....WW.WW.W....W..W.........WW..W.....W.WWW.W....W.W.....
.WW...WWW.W..W.W..W..W.W...W.W..WW..WWW.W.W....WW..WW........W..W.........W.W.......W.WW........W...
..W.W...W...W...W.WW.....W....W.W.............W..W..W.W.W.WW...WW..W.............W....WW............
.....W.W.W.WWW....W.W....W.WW.W..W.WWW..W......WWW....W...W.W.W.....W..W.W.W.....W.WW.WWW..W.......W
WW.W.......W.....W......W...W....W..WW..W.....W..W..W.......W.W..W.W.....W.........W...W..W..W.....W
.WW.....W...W..W.W.........W......W.WW.W...W...W....W..WWW..WW.W....W........WW.W.WW........W.W.....
W...W.....W..W...W.WW...W..WW.W.WW..WW.......W..W.W......W.WW..W....WW.....WWW.WWW..W....WW...WW..W.
W.........WW.W.W.W.....W..W...................W.W......W............W.W..W.W...........W.W.W..W...WW
...W.W...WW...W.W..WW......WW.......W.........W.W.....W.W.....WW...W.W..........W.....W.............
......WW.W.W...W...W.W.W..WW.WW.....W.W..........W.WW..WWW....W.W......W.W..W.....W.W..........W....
......WW....W...WWWW..WW.WW.....W.W.W....W..W.W...W...W.........W.W.W.W........W.....WW.W..W...W....
.......WWW..W......W...W.W.WW.................W.....W.W.......W..W....W..W.W.W..W.W......WW...WW.W..
......W.......W.WWW..W......W...WW........W......W........W.W.W.....W.......WW.....W...W...W.W......
W...W..W..W......WWW.....WW...W...WW..W.W...W.WW...W....W.W....W.W..W....W........W.......W......W..
W..W..W.W.....WW..WWW.........W...W....WW.W.....WW..W..W.W.........W.W..WW................W.W.....WW
...WW..W...W..W........W....WW.WW....W.W.WW.........WW..W....W...W...W.W..W..W.W..W..W.W...W..WWWW..
W.WW.W.W..WWWW......W.WWW...W..........W.W....WW.........WW..W......W.....WWW....W.....W...W...W.W..
...WW..W........W..WW.W....WWW.W.W....WW....WW.W...W......W...........WWW..WW.W..WWW......WW......W.
.W.......WWW..W.W.W.W....WW.WWW.W......W.W.W..W.....W.W.WW..W......W..W.WW.W..W..W......W..W.W..W...
.......W...W.W...........W.W..W....W.W..W...W..WW....W...WW...WW..W..W.WWWWW................W.......
WW..W.....WW..W..W.....................W...WWW.W......W........WW.W.W....W.W.....W.WWW..W.........WW
........W...W....W.......W..WW.W..........W.....W......W....W.W..W..W..W...W.W.......W.W.......W....
.W...WW......W...W.....W...W..WW.W..WW........WW.....WW.........WWWWW...W....WW...............W.....
W.........WWW.W..W....WWW...WWWW..W........W........W..W.WW.....W.WW..W..W...WW.....W.WWWWWWW.W.WW.W
..W.W........WW...........W.W......W.W.W.W.WW.....W..W....W.....W.W..WW.....W...W.....W....W.W.....W
.WW.....W...W...W.W..W..W....W...WW....WW.......W.W...W.....W....W.WWW.W..........WWWW.W..WW........
WW.W.W..W......WW......W....W....WW....WWWW.WW..WW..WW..WWW...W..W..W..........W..W.........W.......
.....W...W..W.......W.........W.W..W..W.WW.WW...W.....W...W...W.......W..W..W.W..W..WW...W...W..W...
W...W.....W.............WW........W.......W.....W...W.......WW..W.W.....W..W..W..W....W.......W...W.
.W......W.WW.....W.W...W..W..W...WW.W...WW.....WWWW.....W.WW........W..W..W...W.W.W.WW..WW...W......
W..WW....WW.W...W...........W.W.W.W..WW................W.WWW.....W..W.W....W.....W.....W.W......W..W
W........WW...W...W..WW.W.W..W..W....WW.W.....W.W.......W.........W.WW.W.....W.........W.WW.....WWWW
...W..W...W...WW.WWW...WWW.WWWWW......WW.W......W....WW...WWWWW.....W.........W.......W....W...WWW..
...WW..WW.WW..W.W.W...WW...W...W.....W.....W....W.....W...WW..W....W..W.W.W..W.....W......W....W..WW
.W..WW....W......W.W.......W.WW.....WW.W....W.W.........WW.W.W..........WWWW.W.WW....WW..W.......WW.
WWWW.......W.W..W......WW..W..W.W.W..WWW....W..W.W..W....W...W...WWW.....W.......WW...W.W.....WWW...`;

console.log(pathFinder(m)); // false

// the solution failed for memmory space


/* best solution

function pathFinder(maze){
  const rows = maze.split(`\n`).map(l => l.split(``));
  const n = rows.length - 1;
  const moveTo = (x, y) => {
    if (x < 0 || y < 0 || x > n || y > n || rows[y][x] !== '.') {
      return false;
    }

    if (x === n && y === n) {
      return true;
    }

    rows[y][x] = `x`;

    return moveTo(x - 1, y)
      || moveTo(x + 1, y)
      || moveTo(x, y - 1)
      || moveTo(x, y + 1);
  }

  return moveTo(0, 0);
}

 */