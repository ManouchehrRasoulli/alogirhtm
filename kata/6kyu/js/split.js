function solution(str) {
    if (str.length % 2 === 1) {
        str += "_";
    }

    let sl = [];
    for (let i = 0; i < str.length; i += 2) {
        sl.push(str[i] + str[i + 1]);
    }

    return sl;
}

console.log(solution("abcdef"), ["ab", "cd", "ef"]);
console.log(solution("abcdefg"), ["ab", "cd", "ef", "g_"]);
console.log(solution(""), []);