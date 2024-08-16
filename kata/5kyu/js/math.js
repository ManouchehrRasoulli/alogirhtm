Math.round = function (number) {
    console.log("Math.round");
    if (isNaN(number))
        return number;

    let v = String(number).split('.');
    if (v.length === 1) {
        return Number(v[0]);
    }

    let p2 = Number(v[1]);
    let base = Math.pow(10, v[1].length - 1);
    if (p2 >= 5 * base) {
        return Number(v[0]) + 1;
    }

    return Number(v[0]);
};

Math.ceil = function (number) {
    console.log("Math.ciel");
    if (isNaN(number))
        return number;

    let v = String(number).split('.');
    if (v.length === 1) {
        return Number(v[0]);
    }

    let p2 = Number(v[1]);
    if (p2 !== 0) {
        return Number(v[0]) + 1;
    }

    return Number(v[0]);
};

Math.floor = function (number) {
    console.log("Math.floor");
    if (isNaN(number))
        return number;

    return Number(String(number).split('.')[0]);
};