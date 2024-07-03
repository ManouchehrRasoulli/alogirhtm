function arr2bin(arr) {
    var sum = 0;
    for (i = 0; i < arr.length; i++) {
        if (typeof arr[i] === 'number') {
            sum = sum + (arr[i]);
        }
    }
    return sum.toString(2);
}

console.log(arr2bin([1, 2, 3, 4]));
console.log(arr2bin([1, 'a', '2', 1]));
console.log(arr2bin([null]));
console.log(arr2bin([true, true, false, 15]));

