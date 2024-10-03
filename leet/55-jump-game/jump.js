function canJump(nums) {
    let l = nums.length - 1;
    let point = l;
    let lastPoint = l;
    for (let i = l; i >= 0; i--) {
        lastPoint = i;
        if (nums[lastPoint] + i >= point) {
            point = lastPoint;
        }
    }
    return point == 0;
}

console.log(canJump([1, 1, 1, 4]));


