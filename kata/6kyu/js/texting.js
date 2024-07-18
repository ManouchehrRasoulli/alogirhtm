const x = {
    ".": "1",
    ",": "11",
    "?": "111",
    "!": "1111",
    "a": "2",
    "b": "22",
    "c": "222",
    "d": "3",
    "e": "33",
    "f": "333",
    "g": "4",
    "h": "44",
    "i": "444",
    "j": "5",
    "k": "55",
    "l": "555",
    "m": "6",
    "n": "66",
    "o": "666",
    "p": "7",
    "q": "77",
    "r": "777",
    "s": "7777",
    "t": "8",
    "u": "88",
    "v": "888",
    "w": "9",
    "x": "99",
    "y": "999",
    "z": "9999",
    "'": "*",
    "-": "**",
    "+": "***",
    "=": "****",
    " ": "0",
    "1": "1-",
    "2": "2-",
    "3": "3-",
    "4": "4-",
    "5": "5-",
    "6": "6-",
    "7": "7-",
    "8": "8-",
    "9": "9-",
    "0": "0-",
    "*": "*-",
    "#": "#-",
};

function resideOnSameButton(a, b) {
    let al = a.toLowerCase();
    let bl = b.toLowerCase();

    let av = x[al];
    let bv = x[bl];

    if (b === av[0]) { // means digit already hold
        return false;
    }

    return av[0] === bv[0];
}

function isInCharacterSet(a) {
    return /[a-z]/.test(a.toLowerCase());
}

function isUpperCase(a) {
    return /[A-Z]/.test(a);
}

function sendMessage(message) {
    console.log(message);

    let touch = "";
    let isInUpperCase = false;

    for (let i = 0; i < message.length; i++) {
        let m = message[i];
        if (isInCharacterSet(m) && !isInUpperCase && isUpperCase(m)) {
            touch += "#";
            isInUpperCase = true;
        } else if (isInCharacterSet(m) && isInUpperCase && !isUpperCase(m)) {
            touch += "#";
            isInUpperCase = false;
        } else if (i !== 0 && resideOnSameButton(m, message[i - 1])) {
            touch += " ";
        }
        touch += x[m.toLowerCase()];
    }

    return touch;
}

console.log(sendMessage("hey")); // 4433999
console.log(sendMessage("HX")); // #4499
console.log(sendMessage("Hx")); // #44#99
console.log(sendMessage("A-z")); // #2**#9999
console.log(sendMessage("1984")); // 1-9-8-4-
console.log(sendMessage("hi")); // 44 444
console.log(sendMessage("Hi")); // #44#444
console.log(sendMessage("HI")); // #44 444
console.log(sendMessage("one two three")) // 666 6633089666084477733 33
