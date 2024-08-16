function nextEviction(items, index, k) {
    let count = 1;
    let i = index;
    while (count <= k) {
        if (count === k && items[i] !== 'dead') {
            return i;
        }

        if (items[i] !== 'dead') {
            count++;
        }

        i++;
        if (i === items.length) {
            i = 0;
        }
    }
}

function josephus(items, k) {
    let deaths = [];
    let index = 0;
    while (deaths.length !== items.length) {
        index = nextEviction(items, index, k);
        deaths.push(items[index]);
        items[index] = 'dead';
    }

    return deaths;
}

console.log(josephus([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 1), [1, 2, 3, 4, 5, 6, 7, 8, 9, 10])
console.log(josephus([1, 2, 3, 4, 5], 2), [2, 4, 1, 5, 3])
console.log(josephus(["C", "o", "d", "e", "W", "a", "r", "s"], 4), ['e', 's', 'W', 'o', 'C', 'd', 'r', 'a'])
console.log(josephus([1, 2, 3, 4, 5, 6, 7], 3), [3, 6, 2, 7, 5, 1, 4])
console.log(josephus([], 3), [])