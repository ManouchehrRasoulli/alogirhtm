function include(a, b) {
    if (a === b) {
        return true;
    }
    if (b === '' && a !== '') {
        return false;
    }
    return a.includes(b);
}

function findUniq(arr) {
    let arm = [];
    for (let i = 0; i < arr.length; i++) {
        arm[i] = arr[i].toLowerCase().replaceAll(' ', '').split('').sort().join('');
    }

    console.log(arm);

    for (let i = 0; i < arm.length; i++) {
        let isPrefix = false;
        console.log("i", i, arm[i]);
        for (let j = 0; j < arm.length; j++) {
            console.log("j", j, arm[j], arm[i].includes(arm[j]));
            if (include(arm[i], arm[j]) && i !== j) {
                isPrefix = true;
                break;
            }
        }
        if (!isPrefix) {
            return arr[i];
        }
    }

    return -1;
}

console.log(findUniq(['Tom Marvolo Riddle', 'I am Lord Voldemort', 'Harry Potter']));
console.log(findUniq(['    ', 'a', ' ']));
