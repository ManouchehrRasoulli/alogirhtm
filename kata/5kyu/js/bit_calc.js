function btn(binary) {
    let s = 0;
    binary = binary.split('').reverse().join('');
    for (let i = 0; i < binary.length; i++) {
        binary[i] === '1' ? s += Math.pow(2, i) : 0;
    }
    return s;
}

function calculate(num1, num2) {
    return btn(num1) + btn(num2);
}

console.log(btn("011"));
console.log(btn("10"));

console.log(calculate("101", "10"));
