function plus13(x) {
    if (!/[a-zA-Z]/.test(x)) {
        return x;
    }

    if (/[A-Z]/.test(x)) {
        return String.fromCharCode((x.charCodeAt(0) - 65 + 13) % 26 + 65);
    }

    return String.fromCharCode((x.charCodeAt(0) - 97 + 13) % 26 + 97);
}

function rot13(message) {
    let msg = "";
    for (let i = 0; i < message.length; i++) {
        msg += plus13(message[i]);
    }
    return msg;
}

console.log(rot13("A"));
console.log(rot13("Z"));
console.log(rot13("a"));
console.log(rot13("z"));
console.log(rot13("Hello"));
console.log(rot13("Xiau!31aS"));
