/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function (s) {
    return s.split(" ").reverse().filter(function (value, index, array) {
        if (value.length !== 0) {
            return value;
        }
    }).join(" ");
};

console.log(reverseWords("the sky is blue"));
console.log(reverseWords("  hello world  "));
console.log(reverseWords("a good   example"));
