/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
var merge = function (intervals) {
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

console.log(merge([[8, 10], [1, 3], [2, 6], [15, 18]]));
console.log(merge([[1, 4], [4, 5]]));
console.log(merge([[1,4],[2,3]]));
