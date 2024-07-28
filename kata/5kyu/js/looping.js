class Node {
    constructor() {
        this.next = null;
    }

    getNext() {
        return this.next;
    }

    setNext(next) {
        this.next = next;
    }
}

function loop_size(node) {
    let nodes = new Set();
    let pointer = node;
    while (!nodes.has(pointer)) {
        nodes.add(pointer);
        pointer = pointer.next;
    }

    let loopCount = 1;
    let traverser = pointer;
    while (traverser.next !== pointer) {
        traverser = traverser.next;
        loopCount++;
    }

    return loopCount;
}

const A = new Node(), B = new Node();
A.setNext(B);
B.setNext(A);

console.log(loop_size(A));

