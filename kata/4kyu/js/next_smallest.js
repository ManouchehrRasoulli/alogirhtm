function swap(arr, i, j) {
    let s = arr[i];
    arr[i] = arr[j];
    arr[j] = s;
}

function thereIsAnyBigger(n, smmallest) {
    let ns = String(smmallest);
    let items = [];
    for (let i = 0; i < ns.length; i++) {
        items[i] = ns[i];
    }

    let len = items.length;

    for (let i = 0; i < len; i++) {
        for (let j = i + 1; j < len; j++) {
            swap(items, i, j);
            let ns = Number(items.join(''));
            if (String(ns).length === len && ns < n && ns > smmallest) {
                return ns;
            }
            swap(items, i, j);
        }
    }

    return -1;
}

function nextSmaller(n) {
    let ns = String(n);
    let items = [];
    for (let i = 0; i < ns.length; i++) {
        items[i] = ns[i];
    }

    let smallest = -1;
    let len = items.length;

    for (let i = 0; i < len; i++) {
        for (let j = i + 1; j < len; j++) {
            swap(items, i, j);
            let ns = Number(items.join(''));
            if (String(ns).length === len && ns < n && ns > smallest) {
                smallest = ns;
            }
            swap(items, i, j);
        }
    }

    while(true) {
        let nextSmallest = thereIsAnyBigger(n, smallest);
        if (nextSmallest === -1) {
            break;
        }

        if (nextSmallest > smallest && nextSmallest < n) {
            smallest = nextSmallest;
        }
    }

    return smallest;
}

console.log(nextSmaller(531));
console.log(nextSmaller(907));
console.log(nextSmaller(135));
console.log(nextSmaller(414));
console.log(nextSmaller(1234567908)); // 1234567890
console.log(nextSmaller(275993364578)); // 275993358764