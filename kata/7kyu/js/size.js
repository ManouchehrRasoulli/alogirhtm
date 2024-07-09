const vX = new Map();
vX.set("x", 2);

const vT = new Map();
vT.set("s", 36);
vT.set("m", 38);
vT.set("l", 40);

function sizeToNumber(size) {
    size = size.toLowerCase();

    let c = 0;
    for (let i = 0; i < size.length - 1; i++) {
        if (vX.get(size[i]) === undefined) {
            return null; // invalid input
        }

        c = vX.get(size[i]) + c;
    }

    if (vT.get(size[size.length - 1]) === undefined || size[size.length - 1] === 'm' && c !== 0) { // invalid input
        return null;
    }

    if (size[size.length - 1] === 's') {
        c = vT.get(size[size.length - 1]) - c;
    } else {
        c = vT.get(size[size.length - 1]) + c;
    }

    return c;
}

console.log(sizeToNumber("s"));
console.log(sizeToNumber("m"));
console.log(sizeToNumber("l"));
console.log(sizeToNumber("xm"));
console.log(sizeToNumber("xxl"));
console.log(sizeToNumber("xxxx"));
console.log(sizeToNumber("xxs"));
console.log(sizeToNumber("xxxm"));
console.log(sizeToNumber("xxms"));
