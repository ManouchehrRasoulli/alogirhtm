function getInteger(x) {
    let v = x - '0';

    if (isNaN(v) || v < 0 || v > 9) {
        return undefined
    }

    return v;
}

function sliceAdder(slice, x) {
    for (let i = 0; i < slice.length; i++) {
        let v = slice[i] + x;
        if (v < 10) {
            x = 0;
            slice[i] = v;
            break;
        }
        v -= 10;
        x = 1;
        slice[i] = v;
    }

    if (x > 0) {
        slice.push(x);
    }

    return slice;
}

function incrementString(strng) {
    let nums = [];
    let lastIndex = 0;
    for (let i = strng.length; i > 0; i--) {
        let n = getInteger(strng[i - 1]);
        if (n === undefined) {
            lastIndex = i;
            break;
        }
        nums.push(n);
    }

    nums = sliceAdder(nums, 1);
    let string = "";

    for (let i = 0; i < lastIndex; i++) {
        string += strng[i];
    }

    for (let i = nums.length; i > 0; i--) {
        string += nums[i-1];
    }

    return string;
}

console.log(incrementString("string001"));
console.log(incrementString("last12"));
console.log(incrementString("las99"));
console.log(incrementString("las0099"));
console.log(incrementString("0099"));
