/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (numbers, target) {
    for (let i = 0; i < numbers.length; i++) {
        let current = numbers[i];
        let other = target - current;
        let otherIndex = binarySearch(numbers, other);
        if (otherIndex !== -1) {
            if (otherIndex === i) {
                return [i + 1, i + 2];
            }
            return [i + 1, otherIndex + 1];
        }
    }

    return [];
};

/**
 * @param {number[]} arr
 * @param {number} target
 * @return {number}
 */
function binarySearch(arr, target) {
    let left = 0;
    let right = arr.length - 1;

    while (left <= right) {
        const mid = Math.floor((left + right) / 2);

        if (arr[mid] === target) {
            return mid;
        } else if (arr[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }

    return -1;
};

console.log(twoSum([2, 7, 11, 15], 9)); // [1, 2]
console.log(twoSum([2, 3, 4], 6)); // [1, 3]
console.log(twoSum([-1, 0], -1)); // [1, 2]
console.log(twoSum([0, 0, 3, 4], 0)); // [1, 2]
console.log(twoSum([-10, -8, -2, 1, 2, 5, 6], 0)); // [3, 5]