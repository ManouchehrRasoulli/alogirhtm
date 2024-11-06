/**
 * @param {number[]} nums
 * @return {number}
 */
var oneCount = function (num) {
    let count = 0;
    while (num !== 0) {
        count += num & 1;
        num >>= 1;
    }
    return count;
}

/**
 * @param {number[]} nums
 * @return {boolean}
 */
var canSortArray = function (nums) {
    let preSetBits = -1;
    let prevMax = -Infinity;
    let curMax = -Infinity;
    let curMin = -Infinity;

    for (let num of nums) {
        let curSetBit = oneCount(num);
        if (curSetBit !== preSetBits) {
            if (prevMax > curMin)
                return false;

            preSetBits = curSetBit;
            prevMax = curMax;
            curMax = num;
            curMin = num;
        } else {
            curMax = Math.max(curMax, num);
            curMin = Math.min(curMin, num);
        }
    }

    return prevMax <= curMin;
};

console.log(canSortArray([8, 4, 2, 30, 15]));
