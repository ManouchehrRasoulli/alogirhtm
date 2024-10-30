var mem = {};

/**
 * @param {number[]} houses
 * @param {number} k
 * @return {number}
 */
var p = function (houses, k) {
    if (k <= 0) {
        return 0;
    }

    if (mem[k] !== undefined) {
        return mem[k];
    }

    if (k === 1) {
        return houses[k - 1];
    }

    mem[k] = Math.max(p(houses, k - 2) + houses[k - 1], p(houses, k - 1));
    return mem[k];
}

/**
 * @param {number[]} houses
 * @return {number}
 */
var rob = function (houses) {
    mem = new Set();
    return p(houses, houses.length);
};

console.log(rob([1, 2, 3, 1])); // 4
console.log(rob([2, 7, 9, 3, 1])); // 12
console.log(rob([2, 1, 1, 2])); // 4
