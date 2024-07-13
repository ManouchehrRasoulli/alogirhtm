function dif(a, b) {
    if (a > b) {
        return a - b;
    }
    return b - a;
}

function solve(arr, n) {
    console.log(arr, n);

    let cats = [];
    let dogs = [];

    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === "C") {
            cats.push(i);
            continue
        }
        dogs.push(i);
    }

    let catchCount = 0;
    for (let catIndex = 0; catIndex < cats.length; catIndex++) {
        for (let dogIndex = 0; dogIndex < dogs.length; dogIndex++) {
            if (dogs[dogIndex] === -1) {
                continue;
            }

            if (dif(dogs[dogIndex], cats[catIndex]) <= n) {
                catchCount++;
                dogs[dogIndex] = -1;
                break;
            }
        }
    }

    return catchCount;
}

console.log(solve(['D', 'C', 'C', 'D', 'C'], 1)); // 2
console.log(solve(['C', 'C', 'D', 'D', 'C', 'D'], 2)); // 3
console.log(solve(['C', 'C', 'D', 'D', 'C', 'D'], 1)); // 2
console.log(solve(['D', 'C', 'D', 'C', 'C', 'D'], 3)); // 3
console.log(solve(['C', 'C', 'C', 'D', 'D'], 3)); // 2
console.log(solve(['C', 'C', 'C', 'D', 'D'], 2)); // 2
console.log(solve(['C', 'C', 'C', 'D', 'D'], 1)); // 1
