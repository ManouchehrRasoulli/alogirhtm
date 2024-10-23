/**
 * @param {number} x
 * @return {number}
 */
var reverse = function (x) {
    if (x < 0) {
        return -reverse(-x);
    }

    let rx = 0;
    while (x !== 0) {
        let v = x % 10;
        x = Math.floor(x / 10);
        rx *= 10;
        rx += v;
    }

    if (rx > 2147483647) {
        return 0;
    }

    return rx;
};

console.log(reverse(123)); // 321
console.log(reverse(-123));
console.log(reverse(120)); // 21
console.log(reverse(1201)); // 1021
