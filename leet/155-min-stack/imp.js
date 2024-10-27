class BetterMinStack {
    constructor() {
        this._array_ = [];
        this._min_stack_ = [];
    };
    push(val) {
        this._array_.push(val);
        if (this._min_stack_.length === 0) {
            this._min_stack_.push(val);
            return;
        }

        let current = this._min_stack_[this._min_stack_.length - 1];
        this._min_stack_.push(Math.min(current, val));
    };

    pop() {
        this._min_stack_.pop();
        return this._array_.pop();
    };

    top() {
        return this._array_[this._array_.length - 1];
    };

    getMin() {
        return this._min_stack_[this._min_stack_.length - 1];
    };
};
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
};

let minStack = new BetterMinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
console.log(minStack.getMin()); // return -3
console.log(minStack.pop());
console.log(minStack.top());    // return 0
console.log(minStack.getMin()); // return -2
