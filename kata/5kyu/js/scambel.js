function strToMp(str) {
    let strMp = new Map();

    for (let i = 0; i < str.length; i++) {
        let v = strMp.get(str.charAt(i));
        if (v === undefined) {
            strMp.set(str.charAt(i), 1);
            continue;
        }
        strMp.set(str.charAt(i), v + 1);
    }

    return strMp;
}

function scramble(str1, str2) {
    console.log(str1, str2);
    console.log(str1.length, str2.length);

    if (str2.length > str1.length) {
        return false;
    }

    let str1mp = strToMp(str1);

    for (let i = 0; i < str2.length; i++) {
        let k = str2.charAt(i);
        let v = str1mp.get(k);
        if (v === undefined || v === 0) {
            return false;
        }

        str1mp.set(k, v - 1);
    }

    return true;
}

console.log(scramble('rkqodlw', 'world'));
console.log(scramble('cedewaraaossoqqyt', 'codewars'));
console.log(scramble('katas', 'steak'));
console.log(scramble('scriptjavx', 'javascript'));
