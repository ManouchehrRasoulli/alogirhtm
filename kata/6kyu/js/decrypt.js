function charToAscii(char) {
    let code = char.charCodeAt(0);
    if (code < 97 || code > 122) { // between a-z
        return undefined;
    }
    return code - 97;
}

function decrypt(encryption) {
    let m = new Map();
    for (let i = 0; i < encryption.length; i++) {
        let code = charToAscii(encryption[i]);
        if (code === undefined) { // ignore
            continue;
        }
        let v = m.get(code);
        m.set(code, isNaN(v) ? 1 : v + 1);
    }

    const fill = (n, val = 0) => Array(n).fill(val);
    let arr = fill(26, 0);

    m.forEach(function (value, key, map) {
        arr[key] = value;
    });

    let str = "";
    for (let i = 0; i < arr.length; i++) {
        str += arr[i];
    }

    return str;
}

console.log(decrypt("Asbauczaa1?!@8"));
console.log(decrypt("@c{yxum'pqjuv)>`f_]bcd+v`$ix(fn[@&cc%:/wo@sks3r%ds4bm-i,x_j>]u.2yl_bux3;=o:stgd(ip&2-shi149>3rne:_(ntik4zkjv)z!j(9&njwk-{0l`zmrj:mpej<#r6fb$(;-ry-^94zwg5!z5\\8n?+ntwurb4]ehju.5uq5ukh+memf%y\\pc(le%`3m@[)%c;\"#y\"`?6'6be]@v&subo4#k.'+%zte-s.)x{mhu-^agl!piu816j:+/nogo;tncq)mu0`]9![88zq%u\"efo+nis9sh[bg[[bf"));
