const right = 'r';
const down = 'd';
const left = 'l';
const up = 'u';


// step

function generateSequence(n, step) {
    let sequence = [];
    if (n === 0) {
        return sequence;
    }

    if (step === 0) {
        sequence.push({
            count: n - 1,
            direction: right,
        });
        let sq = generateSequence(n - 1, step + 1);
        sequence.push(...sq);
        return sequence;
    }

    if (step % 2 === 0) {
        sequence.push({
            count: n,
            direction: up,
        }, {
            count: n,
            direction: right,
        });
    } else {
        sequence.push({
            count: n,
            direction: down,
        }, {
            count: n,
            direction: left,
        });
    }

    let sq = generateSequence(n - 1, step + 1);
    sequence.push(...sq);
    return sequence;
}

function generateLocations(loc, count, direction) {
    let locations = [];
    let x = loc[0];
    let y = loc[1];

    for (let i = 0; i < count; i++) {
        switch (direction) {
            case right:
                y += 1;
                break;
            case down:
                x += 1;
                break;
            case left:
                y -= 1;
                break;
            case up:
                x -= 1;
                break;
        }
        locations.push([x, y]);
    }

    return locations;
}

function step(n) {
    let locations = [[0, 0]];
    let count = 0;
    let sequence = generateSequence(n, 0);
    for (let i = 0; i < sequence.length; i++) {
        let loc = generateLocations(locations[count], sequence[i].count, sequence[i].direction);
        locations.push(...loc);
        count = locations.length - 1;
    }
    return locations;
}

snail = function (array) {
    let steps = step(array.length);
    let items = [];
    for (let i = 0; i < steps.length; i++) {
        items.push(array[steps[i][0]][steps[i][1]]);
    }
    if (items.length === 1 && items[0] === undefined) {
        return [];
    }
    return items;
};
