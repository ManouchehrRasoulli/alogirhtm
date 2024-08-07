function permutations(string) {
    let result = [];
    if (string.length === 0) return [];
    if (string.length === 1) return [string];

    for (let i = 0; i < string.length; i++) {
        const currentNum = string[i]
        const remainingNums = string.slice(0, i).concat(string.slice(i + 1))
        const remainingNumsPermuted = permutations(remainingNums)
        for (let j = 0; j < remainingNumsPermuted.length; j++) {
            const permutedArray = [currentNum].concat(remainingNumsPermuted[j])
            let x = permutedArray.join('');
            if (result.indexOf(x) === -1) {
                result.push(x);
            }
        }
    }

    return result;
}

console.log(permutations('a'));
console.log(permutations('aabb'));