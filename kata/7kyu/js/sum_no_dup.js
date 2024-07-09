function sumNoDuplicates(numList) {
    let m = new Map();
    for (let i = 0; i < numList.length; i++) {
        let v = m.get(numList[i]);
        if (v === undefined) {
            m.set(numList[i], 1);
            continue;
        }
        m.set(numList[i], v + 1);
    }

    let sum = 0;
    m.forEach(function (value, key, map) {
        if (value === 1) {
            sum += key;
        }
    })

    return sum;
}

console.log(sumNoDuplicates([12, 12, 33, 1]));
console.log(sumNoDuplicates([1, 1, 2, 3]));
