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
    let m = new Map();

    for (let i = 0; i < nums.length; i++) {
        let v = m.get(nums[i]);
        if (v === undefined) {
            m.set(nums[i], 1);
            continue;
        }

        m.set(nums[i], m.get(nums[i]) + 1);
    }

    console.log(m);

    let r = [nums.length];
    for (let i = 0; i < nums.length; i++) {
        let count = 0;
        for (let j = 0; j < nums[i]; j++) {
            let v = m.get(j);
            if (v !== undefined) {
                count += v;
            }
        }
        r[i] = count;
    }

    return r;
}

console.log(smaller([5, 4, 3, 2, 1]));
console.log(smaller_b([5, 4, 3, 2, 1]));
