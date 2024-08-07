function addWitCary(a, b, c) {
    let x = Number(a) + Number(b) + Number(c);
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

    return x.split('').reverse().join('');
}

console.log(add('63829983432984289347293874', '90938498237058927340892374089'));
// 91002328220491911630239667963

