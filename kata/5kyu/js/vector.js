Vector = class {
    __cmp__
    __ln__

    constructor(items) {
        this.__cmp__ = items;
        this.__ln__ = items.length;
    };

    add(v) {
        if (this.__ln__ !== v.__ln__) {
            throw new Error("mismatched-len-add");
        }
        let x = [];
        for (let i = 0; i < this.__ln__; i++) {
            x[i] = this.__cmp__[i] + v.__cmp__[i];
        }
        return new Vector(x);
    };

    subtract(v) {
        if (this.__ln__ !== v.__ln__) {
            throw new Error("mismatched-len-subtract");
        }
        let x = [];
        for (let i = 0; i < this.__ln__; i++) {
            x[i] = this.__cmp__[i] - v.__cmp__[i];
        }
        return new Vector(x);
    };

    dot(v) {
        if (this.__ln__ !== v.__ln__) {
            throw new Error("mismatched-len-dot");
        }
        let x = 0;
        for (let i = 0; i < this.__ln__; i++) {
            x += this.__cmp__[i] * v.__cmp__[i];
        }
        return x;
    };

    norm() {
        let x = 0;
        for (let i = 0; i < this.__ln__; i++) {
            x += Math.pow(this.__cmp__[i], 2);
        }
        return Math.sqrt(x);
    };

    equals(v) {
        if (this.__ln__ !== v.__ln__) {
            return false;
        }

        for (let i = 0; i < this.__ln__; i++) {
            if (this.__cmp__[i] !== v.__cmp__[i]) {
                return false;
            }
        }

        return true;
    };


    toString() {
        return "(" + this.__cmp__.join(',') + ")";
    };
};

let a = new Vector([1, 2]);
let b = new Vector([3, 4]);
console.log(a.add(b));
console.log(a.toString());
