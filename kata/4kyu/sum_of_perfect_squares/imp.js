var dp = [0]

function sumOfSquares(n) {
    if (n < dp.length) {
        return dp[n];
    }

    let start = dp.length;
    let end = n + 1;

    for (let i = start; i <= end; i++) {
        dp[i] = n + 1;
        for (let j = 1; j * j <= i; j++) {
            dp[i] = Math.min(dp[i], dp[i - j * j] + 1);
        }
    }

    return dp[n];
}

console.log(sumOfSquares(13)); // 2
console.log(sumOfSquares(12)); // 3

