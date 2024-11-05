/**
 * @param {string} s
 * @return {number}
 */
var minChanges = function (s) {
    let changes = 0;
    for (let i = 1; i < s.length; i += 2) {
        if (s[i] !== s[i - 1]) {
            changes++;
        }
    }

    return changes;
};

console.log(minChanges("1001"));
console.log(minChanges("01"));
console.log(minChanges("0000"));

