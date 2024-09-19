var screen = new Map();

// A
var aMoves = new Map();
aMoves.set('B', []);
aMoves.set('C', ['B']);
aMoves.set('D', []);
aMoves.set('E', []);
aMoves.set('F', []);
aMoves.set('G', ['D']);
aMoves.set('H', []);
aMoves.set('I', ['E']);
screen.set('A', aMoves);

// B
var bMoves = new Map();
bMoves.set('A', []);
bMoves.set('C', []);
bMoves.set('D', []);
bMoves.set('E', []);
bMoves.set('F', []);
bMoves.set('G', []);
bMoves.set('H', ['E']);
bMoves.set('I', []);
screen.set('B', bMoves);

// C
var cMoves = new Map();
cMoves.set('A', ['B']);
cMoves.set('B', []);
cMoves.set('D', []);
cMoves.set('E', []);
cMoves.set('F', []);
cMoves.set('G', ['E']);
cMoves.set('H', []);
cMoves.set('I', ['F']);
screen.set('C', cMoves);

// D
var dMoves = new Map();
dMoves.set('A', []);
dMoves.set('B', []);
dMoves.set('C', []);
dMoves.set('E', []);
dMoves.set('F', ['E']);
dMoves.set('G', []);
dMoves.set('H', []);
dMoves.set('I', []);
screen.set('D', dMoves);

// E
var eMoves = new Map();
eMoves.set('A', []);
eMoves.set('B', []);
eMoves.set('C', []);
eMoves.set('D', []);
eMoves.set('F', []);
eMoves.set('G', []);
eMoves.set('H', []);
eMoves.set('I', []);
screen.set('E', eMoves);

// F
var fMoves = new Map();
fMoves.set('A', []);
fMoves.set('B', []);
fMoves.set('C', []);
fMoves.set('D', ['E']);
fMoves.set('E', []);
fMoves.set('G', []);
fMoves.set('H', []);
fMoves.set('I', []);
screen.set('F', fMoves);

// G
var gMoves = new Map();
gMoves.set('A', ['D']);
gMoves.set('B', []);
gMoves.set('C', ['E']);
gMoves.set('D', []);
gMoves.set('E', []);
gMoves.set('F', []);
gMoves.set('H', []);
gMoves.set('I', ['H']);
screen.set('G', gMoves);

// H
var hMoves = new Map();
hMoves.set('A', []);
hMoves.set('B', ['E']);
hMoves.set('C', []);
hMoves.set('D', []);
hMoves.set('E', []);
hMoves.set('F', []);
hMoves.set('G', []);
hMoves.set('I', []);
screen.set('H', hMoves);

// I
var iMoves = new Map();
iMoves.set('A', ['E']);
iMoves.set('B', []);
iMoves.set('C', ['F']);
iMoves.set('D', []);
iMoves.set('E', []);
iMoves.set('F', []);
iMoves.set('G', ['H']);
iMoves.set('H', []);
screen.set('I', iMoves);

function countPatternsFrom(firstPoint, length) {
    if (length < 1 && length > 9) {
        return 0;
    }

    let q = [];
    q.push(firstPoint);
    let count = 0;

    while (q.length !== 0) {
        let item = q.shift();
        if (item.length === length) {
            count += 1;
            continue;
        }
        let lastChar = item[item.length - 1];
        let moves = screen.get(lastChar);

        let iter = moves.keys();
        let next = iter.next();
        while (!next.done) {
            let mv = next.value;
            let touched = moves.get(mv);
            // 1 - v is empty then check if key is not in item
            if (touched.length === 0 && item.indexOf(mv) === -1) {
                let newItem = item + mv;
                q.push(newItem);
            }
            // 2 - v is not empty then check if item already in item
            else if (item.indexOf(touched[0]) !== -1 && item.indexOf(mv) === -1) {
                let newItem = item + mv;
                q.push(newItem);
            }
            next = iter.next();
        }
    }

    return count;
}

