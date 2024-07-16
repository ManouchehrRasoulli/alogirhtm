function onlyDuplicates(str) {
    let d = new Map();

    for (let i = 0; i < str.length; i++) {
        (d[str[i]] === undefined) ? d[str[i]] = 1 : d[str[i]]++;
    }

    let x = "";

    for (let i = 0; i < str.length; i++) {
        if (d[str[i]] > 1) {
            x += str[i];
        }
    }

    return x;
}

console.log(onlyDuplicates('abccdefee'), 'cceee', "onlyDuplicates('abccdefee')")
console.log(onlyDuplicates('hello'), 'll', "onlyDuplicates('hello')")
console.log(onlyDuplicates('foundersandcoders'), 'ondersndoders', "onlyDuplicates('foundersandcoders')")
console.log(onlyDuplicates('colloquial'), 'ollol', "onlyDuplicates('colloquial')")
