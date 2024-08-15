let mem = [0, 0, 1, 2];

function fibonacci(n) {
    if (n < 2) {
        return n;
    }
    if (n < mem.length) {
        return mem[n];
    }
    let x = fibonacci(n - 1) + fibonacci(n - 2);
    mem.push(x)
    return mem[n];
}

console.log(fibonacci(3), "should be 2");
console.log(fibonacci(4), "should be 3");
console.log(fibonacci(5), "should be 5");
console.log(fibonacci(10), "should be 55");
console.log(fibonacci(15), "should be 610");
console.log(fibonacci(20), "should be 6765");
