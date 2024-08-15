/*
function totalSum(n) {
    let sum = 0;
    for (let i = 0; i <= n; i++) {
        sum += i;
    }
    return sum;
}

function removeNb(n) {
    console.log(n);
    let excluded = [];

    let sum = totalSum(n);

    for (let i = 0; i <= n; i++) {
        for (let j = 0; j <= n; j++) {
            if ((i * j) === (sum - (i + j))) {
                excluded.push([i, j]);
            }
        }
    }

    return excluded;
}
 */

function removeNb(n) {
    // from the instruction:
    // a * b = S(n) - a - b

    // and the sum of all first n natural numbers is
    // S(n) = n * (n + 1) / 2

    // so we can refrase the first formula like this:
    // a * b = n * (n + 1) / 2 - a - b
    // a * b + b = n * (n + 1) / 2 - a
    // b * (a + 1) = n * (n + 1) / 2 - a
    // b = (n * (n + 1) / 2 - a) / (a + 1)

    // but a and b are natural and up to n

    var results = [];
    for (var a = 1; a <= n; a++) {
        var b = (n * (n + 1) / 2 - a) / (a + 1);
        if (b % 1 === 0 && b <= n) results.push([a, b]);
    }
    return results;
}

console.log(removeNb(26)) // output [[15,21], [21,15]]
console.log(removeNb(100)) // output []
console.log(removeNb(1000003))