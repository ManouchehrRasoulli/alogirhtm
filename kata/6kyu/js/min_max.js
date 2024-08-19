function flatten() {
    let flat = [];
    for (let i = 0; i < arguments.length; i++) {
        if (Array.isArray(arguments[i])) {
            let fx = flatten(...arguments[i]);
            for (let j = 0; j < fx.length; j++) {
                flat.push(fx[j]);
            }
            continue;
        }
        flat.push(arguments[i]);
    }
    return flat;
}

function max(){
    if (arguments.length === 0) {
        return 0;
    }
    let mx = -1000000000;
    let a = flatten(...arguments);
    for (let i = 0; i < a.length; i++) {
        let iNumber = Number(a[i]);
        if (isNaN(iNumber)) {
            return NaN;
        }
        if (iNumber > mx) {
            mx = iNumber;
        }
    }
    if (mx === -1000000000) {
        return 0;
    }
    return mx;
}

function min(){
    if (arguments.length === 0) {
        return 0;
    }
    let mn = 1000000000;
    let a = flatten(...arguments);
    for (let i = 0; i < a.length; i++) {
        let iNumber = Number(a[i]);
        if (isNaN(iNumber)) {
            return NaN;
        }
        if (iNumber < mn) {
            mn = iNumber;
        }
    }
    if (mn === 1000000000) {
        return 0;
    }
    return mn;
}

console.log(max([1, 2, 3, [4]]));
