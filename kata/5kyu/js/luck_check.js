function charToAscii(char) {
    return char.charCodeAt(0);
}

function luckCheck(ticket) {
    console.log(ticket);

    if (/[a-zA-Z]/.test(ticket)) {
        throw new Error("Parameter is not a number!");
    }

    let firstHalf = 0;
    let secondHalf = 0;

    for (let i = 0; i < ticket.length; i++) {
        let code = charToAscii(ticket[i]) - charToAscii('0');
        if (i < ticket.length / 2) {
            firstHalf += code;
        } else {
            secondHalf += code;
        }
    }

    return firstHalf === secondHalf;
}

console.log(luckCheck('683179')); // true
console.log(luckCheck('683000')); // false
console.log(luckCheck('56328116')); // true
console.log(luckCheck('003111'))// true
console.log(luckCheck('1234')) // false
console.log(luckCheck('6F43E8')) // throw error
