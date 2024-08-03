function findUniq(arr) {
    let s = new Map();

    for (let i = 0; i < arr.length; i++) {
        let v = s.get(arr[i]);
        if (v === undefined) {
            s.set(arr[i], 1);
            continue;
        }
        s.set(arr[i], v + 1);
    }

    for (let i = 0; i < arr.length; i++) {
        let v = s.get(arr[i]);
        if (v === 1) {
            return arr[i];
        }
    }

    return -1;
}
