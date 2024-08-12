class Node {
    constructor(value, left = null, right = null) {
        this.value = value;
        this.left = left;
        this.right = right;
    }
}

function treeByLevels(rootNode) {
    if (rootNode === null) {
        return [];
    }

    let levelOrder = [];

    let q = [];
    q.push(rootNode);

    while (q.length !== 0) {
        let node = q.shift();
        levelOrder.push(node.value);
        if (node.left !== null) {
            q.push(node.left);
        }
        if (node.right !== null) {
            q.push(node.right);
        }
    }

    return levelOrder;
}

const treeOne =
    new Node(2,
        new Node(8,
            new Node(1),
            new Node(3)
        ),
        new Node(9,
            new Node(4),
            new Node(5)
        )
    );
console.log(treeByLevels(treeOne));
