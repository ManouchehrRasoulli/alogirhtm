/**
 * @param {number[]} nums
 * @return {number[]}
 */
var largestDivisibleSubset = function (nums) {
    if (nums.length === 0)
        return [];

    nums.sort((a, b) => a - b);
    let dp = Array(nums.length).fill().map(() => []);
    dp[0] = [nums[0]];

    let maxSubset = dp[0];

    for (let i = 1; i < nums.length; i++) {
        let maxSubSetI = [];
        for (let j = 0; j < i; j++) {
            if (nums[i] % nums[j] === 0 && dp[j].length > maxSubSetI.length) {
                maxSubSetI = dp[j];
            }
        }

        dp[i] = [...maxSubSetI, nums[i]];

        if (dp[i].length > maxSubset.length) {
            maxSubset = dp[i];
        }
    }

    return maxSubset;
};

console.log(largestDivisibleSubset([1, 4, 2, 8]));
console.log(largestDivisibleSubset([1, 2, 3]));
