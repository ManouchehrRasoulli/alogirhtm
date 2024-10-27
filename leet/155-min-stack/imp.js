
class MinStack {
    constructor() {
        this._array_ = [];
    };
    push(val) {
        this._array_.push(val);
    };

    pop() {
        return this._array_.pop();
    };

    top() {
        return this._array_[this._array_.length - 1];
    };

    getMin() {
        return Math.min(...this._array_);
    };
}


let minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
console.log(minStack.getMin()); // return -3
console.log(minStack.pop());
console.log(minStack.top());    // return 0
console.log(minStack.getMin()); // return -2
