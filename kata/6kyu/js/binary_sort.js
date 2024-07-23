function oneCount(binary) {
    let oneCount = 0;
    for (let i = 0; i < binary.length; i++) {
        if (binary[i] === '1') {
            oneCount++;
        }
    }

    return oneCount;
}

function sortByBit(arr) {
    arr.sort((a, b) => {
        let ac = oneCount(a.toString(2));
        let bc = oneCount(b.toString(2));
        if (ac === bc) {
            return a - b;
        }
        return ac - bc;
    });
    return arr;
}

let ar = [3, 8, 3, 6, 5, 7, 9, 1];
console.log(sortByBit(ar));
