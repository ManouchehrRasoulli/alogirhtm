// smaller
// O(n^2)
function smaller(nums) {
    let r = [nums.length];

    for (let i = 0; i < nums.length; i++) {
        let count = 0;
        for (let j = i + 1; j < nums.length; j++) {
            if (nums[j] < nums[i]) {
                count++;
            }
        }
        r[i] = count;
    }

    return r;
}

// smaller_b
// O(n)
function smaller_b(nums) {
    
}

console.log(smaller([5, 4, 3, 2, 1]));
