/**
 * @param {number} start
 * @param {number} end
 * @param {number} current
 * @param {number} target
 * @param {number[]} list
 * @return {number}
 */
var binarySearch = function (start, end, current, targe, list) {
    if (start >= end) {
        return NaN;
    }

    let mid = Math.floor((end + start) / 2);
    let v = current + list[mid];
    if (v === targe) {
        return list[mid];
    }

    if (v < targe) {
        return binarySearch(mid + 1, end, current, targe, list);
    }

    return binarySearch(start, mid, current, targe, list);
}

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
    let sorted = [];
    sorted.push(...nums);
    sorted.sort(function (a, b) { return a - b });

    let resault = [];
    for (let i = 0; i < sorted.length; i++) {
        let v = sorted[i];
        let other = binarySearch(i + 1, nums.length, v, target, sorted);

        if (!isNaN(other)) {
            resault = [v, other];
            break;
        }
    }

    if (resault.length !== 0) {
        let a = nums.indexOf(resault[0])
        nums.splice(a, 1);
        let b = nums.indexOf(resault[1])
        if (b >= a) {
            b += 1;
        }
        let v = [a, b];
        return v.sort(function (a, b) { return a - b });
    }

    return [];
};

console.log(twoSum([3, 2, 3], 6));
console.log(twoSum([-1, -2, -3, -4, -5], -8));
console.log(twoSum([0, 4, 3, 0], 0));
console.log(twoSum([0, 3, -3, 4, -1], -1));


