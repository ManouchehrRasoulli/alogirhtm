function addWitCary(a, b, c) {
    let x = Number(a) + Number(b) + Number(c);
    if (x < 10) {
        return [String(0), String(x)];
    }
    return [String(Math.floor(x / 10)), String(x % 10)];
}

function addOnlyCarry(a, c) {
    let x = Number(a) + Number(c);
    if (x < 10) {
        return [String(0), String(x)];
    }
    return [String(Math.floor(x / 10)), String(x % 10)];
}

function add(a, b) {
    let x = '';

    let an = a.length - 1;
    let bn = b.length - 1;
    let c = '0';
    let av = '0';
    let bv = '0';

    while (an >= 0 || bn >= 0) {
        (an >= 0) ? av = a[an] : av = '0';
        (bn >= 0) ? bv = b[bn] : bv = '0';
        let f = addWitCary(av, bv, c);
        c = f[0];
        x += f[1];
        an -= 1;
        bn -= 1;
    }

    if (c !== '0') {
        x += c;
    }

    return x.split('').reverse().join('').replace(/^0+/, '');
}

function multiplyOneByOne(a, b) {
    let x = Number(a) * Number(b);
    if (x < 10) {
        return [String(0), String(x)];
    }
    return [String(Math.floor(x / 10)), String(x % 10)];
}

function multiplyByOne(a, b) {
    let x = '';

    let an = a.length - 1;
    let c = '0';
    let av = '0';

    while (an >= 0) {
        (an >= 0) ? av = a[an] : av = '0';
        let f = multiplyOneByOne(av, b);
        let ff = addOnlyCarry(f[1], c);
        c = String(Number(ff[0]) + Number(f[0]));
        x += ff[1];
        an--;
    }

    x += c;

    return x.split('').reverse().join('').replace(/^0+/, '');
}

function multiply(a, b) {
    let x = '0';

    a = a.replace(/^0+/, '');
    b = b.replace(/^0+/, '');

    if (a.length === 0 || b.length === 0) {
        return x;
    }

    let bn = b.length - 1;
    let bv = '0';

    while (bn >= 0) {
        let v = multiplyByOne(a, b[bn]);
        for (let i = bn; i < b.length - 1; i++) {
            v += "0";
        }
        x = add(x, v);
        bn--;
    }

    return x;
}

console.log(multiply("2", "10"));

