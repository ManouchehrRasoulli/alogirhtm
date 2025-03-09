function duplicateCount(text){
    let dup = new Map();

    text = text.toLowerCase();
    for (let char of text) {
        let v = dup.get(char);
        if (v === undefined) {
            dup.set(char, 1);
            continue
        }

        dup.set(char, v+1);
    }

    let count = 0;
    dup.forEach(function (value, key, map) {
        if (value > 1) {
            count++;
            console.log(value, key);
        }
    })

    return count;
}

console.log(duplicateCount("some-test-characters"));
