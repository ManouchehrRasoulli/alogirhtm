function toBase16(x) {
    if (x < 0) {
        return "00";
    }
    if (x >= 255) {
        return "FF";
    }

    let v = x.toString(16);
    if (v.length === 1) {
        v = "0" + v;
    }
    return v.toUpperCase();
}

function rgb(r, g, b) {
    return (toBase16(r) + toBase16(g) + toBase16(b));
}

console.log(rgb(0, 0, 0));
console.log(rgb(173, 255, 47));
