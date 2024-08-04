function flatten() {
    console.log(arguments, arguments.length);
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

console.log(flatten(1, null, 5, 'c', [[1], [[[[3], 2]]]]));


