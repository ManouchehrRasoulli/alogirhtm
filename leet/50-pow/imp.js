/**
 * @param {number} x
 * @param {number} n
 * @return {number}
 */
var myPow = function (x, n) {
    console.log(n);

    if (n === 0) {
        return 1;
    }

    if (n < 0) {
        return 1 / myPow(x, -n);
    }

    if (n % 2 === 0) {
        let halth = myPow(x, n / 2);
        return halth * halth;
    }

    return x * myPow(x, n - 1);
};

console.log(myPow(2.00000, 10));
