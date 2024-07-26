function zero(fn) {
    if (typeof fn !== "function") {
        return 0;
    }
    return fn(0);
}

function one(fn) {
    if (typeof fn !== "function") {
        return 1;
    }
    return fn(1);
}

function two(fn) {
    if (typeof fn !== "function") {
        return 2;
    }
    return fn(2);
}

function three(fn) {
    if (typeof fn !== "function") {
        return 3;
    }
    return fn(3);
}

function four(fn) {
    if (typeof fn !== "function") {
        return 4;
    }
    return fn(4);
}

function five(fn) {
    if (typeof fn !== "function") {
        return 5;
    }
    return fn(5);
}

function six(fn) {
    if (typeof fn !== "function") {
        return 6;
    }
    return fn(6);
}

function seven(fn) {
    if (typeof fn !== "function") {
        return 7;
    }
    return fn(7);
}

function eight(fn) {
    if (typeof fn !== "function") {
        return 8;
    }
    return fn(8);
}

function nine(fn) {
    if (typeof fn !== "function") {
        return 9;
    }
    return fn(9);
}

// calculus

function plus(y) {
    return function (x) {
        return x + y;
    };
}

function minus(y) {
    return function (x) {
        return x - y;
    };
}

function times(y) {
    return function (x) {
        return x * y;
    };
}

function dividedBy(y) {
    return function (x) {
        return Math.floor(x / y);
    };
}

console.log(nine(plus(one())));
console.log(seven(minus(nine(plus(one())))));
console.log(six(dividedBy(two())));
console.log(seven(dividedBy(two())));
