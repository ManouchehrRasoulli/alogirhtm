function fibonacciSequence() {
    let current = 1n;
    let next = 1n;
    let iteration = 0;
    const rangeIterator = {
        next() {
            if (iteration < 2) {
                iteration++;
                return {value: 1n, done: false};
            }

            let res = {value: current + next, done: false};
            current = next;
            next = res.value;
            return res;
        }
    }

    return rangeIterator;
}

let iter = fibonacciSequence();

console.log(iter.next());
console.log(iter.next());
console.log(iter.next());
