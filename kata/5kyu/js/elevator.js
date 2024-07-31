function elevator(left, right, {floor: f1, direction: d1}, {floor: f2, direction: d2}) {
    const l1 = Math.abs(f1 - left)
    const l2 = Math.abs(f2 - left)
    const r1 = Math.abs(f1 - right)
    const r2 = Math.abs(f2 - right)
    if (d1 === d2) {
        console.log({left, right, f1, f2, d1, d2})
        if (r1 <= l1 && r2 <= l2) {
            console.log('r test')
            if (Math.sign(f1 - right) === Math.sign(f2 - right) || !(f1 - right) || !(f2 - right)) {
                // right closer
                return 'right right'
            }
        }
        if (l1 <= r1 && l2 <= r2) {
            console.log('l test')
            if (Math.sign(f1 - left) === Math.sign(f2 - left) || !(f1 - left) || !(f2 - left)) {
                // left closer
                return 'left left'
            }
        }
    }

    if (l1 < r1) {
        return 'left right'
    }
    return 'right left'
}

console.log(elevator(0, 3, {floor: 3, direction: 'down'}, {floor: 0, direction: 'up'})); // 'right left'
console.log(elevator(2, 2, {floor: 2, direction: 'down'}, {floor: 3, direction: 'up'})); // 'right left'
console.log(elevator(1, 1, {floor: 1, direction: 'down'}, {floor: 1, direction: 'up'})) // 'right left'
console.log(elevator(4, 3, {floor: 3, direction: 'down'}, {floor: 2, direction: 'down'})) // right right