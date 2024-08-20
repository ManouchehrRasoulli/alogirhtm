function VigenèreCipher(key, abc) {
    this.key = key;
    this.abc = abc;

    this.encode = function (str) {
        let v = "";
        for (let i = 0; i < str.length; i++) {
            let s = this.abc.indexOf(str[i]);
            if (s === -1) {
                v += str[i];
                continue;
            }
            let k = this.key[i % key.length];
            s = s + this.abc.indexOf(k);
            s = s % this.abc.length;
            v += this.abc[s];
        }
        return v;
    };

    this.decode = function (str) {
        let v = "";
        for (let i = 0; i < str.length; i++) {
            let s = this.abc.indexOf(str[i]);
            if (s === -1) {
                v += str[i];
                continue;
            }
            let k = this.key[i % key.length];
            s = s - this.abc.indexOf(k);
            if (s < 0) {
                s += this.abc.length;
            }
            s = s % this.abc.length;
            v += this.abc[s];
        }
        return v;
    };
}

let abc = "abcdefghijklmnopqrstuvwxyz";
let key = "password"
let c = new VigenèreCipher(key, abc);

console.log(c.encode('codewars'), 'rovwsoiv');
console.log(c.decode('rovwsoiv'), 'codewars');

console.log(c.encode('waffles'), 'laxxhsj');
console.log(c.decode('laxxhsj'), 'waffles');

console.log(c.encode('CODEWARS'), 'CODEWARS');
console.log(c.decode('CODEWARS'), 'CODEWARS');
