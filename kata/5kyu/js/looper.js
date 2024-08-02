function makeLooper(str) {
    let index = 0;
    return function () {
        let x = str[index];
        index = (index+1) % str.length;
        return x;
    };
}

let looper = makeLooper("abc");

console.log(looper());
console.log(looper());
console.log(looper());
console.log(looper());
console.log(looper());
