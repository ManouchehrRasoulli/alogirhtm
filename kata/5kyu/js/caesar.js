var CaesarCipher = function (shift) {
    return {
        _sh_(input, shift) {
            let finalString = "";
            input = input.toUpperCase();
            for (let i = 0; i < input.length; i++) {
                let ce = input.charCodeAt(i);
                if (ce < 65 || ce > 90) {
                    finalString += String.fromCharCode(ce);
                    continue;
                }
                let codeShift = ce - 65 + shift;
                if (codeShift < 0) {
                    codeShift += 26;
                }
                let code = (codeShift % 26) + 65;
                finalString += String.fromCharCode(code);
            }
            return finalString;
        },
        encode(input) {
            console.log("encode", shift, input);
            return this._sh_(input, shift);
        },
        decode(input) {
            console.log("decode", shift, input);
            return this._sh_(input, -shift);
        },
    }
};

const c = new CaesarCipher(5);

console.log(c.encode('Z'), 'E');
console.log(c.encode("W"), "B");
console.log(c.encode("A"), "F");
console.log(c.encode("X"), "C");
console.log(c.encode('Codewars'), 'HTIJBFWX');

console.log("decode----");
console.log(c.decode('E'), 'Z');
console.log(c.decode("F"), "A");
console.log(c.decode('HTIJBFWX'), 'CODEWARS');
