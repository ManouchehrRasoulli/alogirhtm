/**
 * @param {number[][]} intervals
 * @param {number[]} newInterval
 * @return {number[][]}
 */
var insert = function (intervals, newInterval) {
    intervals.push(newInterval);
    let removeOverlap = [];
    intervals.sort(function (a, b) {
        if (a[0] < b[0]) {
            return -1;
        }
        return 1;
    });

    removeOverlap.push(intervals[0]);
    for (let i = 1; i < intervals.length; i++) {
        if (intervals[i][0] > removeOverlap[removeOverlap.length - 1][1]) {
            removeOverlap.push(intervals[i]);
            continue;
        }

        if (intervals[i][1] < removeOverlap[removeOverlap.length - 1][1]) {
            continue;
        }

        removeOverlap[removeOverlap.length-1][1] = intervals[i][1];
    }

    return removeOverlap;
};

console.log(insert([[0, 5], [9, 12]], [7, 16]));
console.log(insert([[0, 2], [3, 3], [6, 11]], [9, 15]));
console.log(insert([[1, 5]], [0, 0]));
console.log(insert([[1, 5]], [6, 8]));
console.log(insert([[1, 5]], [2, 7]));
console.log(insert([[2, 7]], [1, 5]));
console.log(insert([[1, 3], [6, 9]], [2, 5]));
console.log(insert([[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8]));
