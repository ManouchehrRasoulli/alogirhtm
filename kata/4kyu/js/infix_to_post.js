function prec(c) {
    if (c === '^')
        return 3;
    if (c === '/' || c === '*')
        return 2;
    if (c === '+' || c === '-')
        return 1;

    return -1;
}

function toPostfix(s) {
    let stack = [];
    let result = "";

    for (let i = 0; i < s.length; i++) {
        if (/[0-9]/.test(s[i])) {
            result += s[i];
            continue;
        }

        if (stack.length === 0) {
            stack.push(s[i]);
            continue;
        }

        if (s[i] === '(') {
            stack.push(s[i]);
            continue;
        }

        if (s[i] === ')') {
            while (stack[stack.length - 1] !== '(') {
                result += stack.pop();
            }
            stack.pop(); // pop the parentheses
            continue;
        }

        if (s[i] === '^') {
            if (stack[stack.length - 1] === s[i]) {
                stack.push(s[i]);
                continue;
            }
        }

        while (stack.length !== 0 && prec(s[i]) <= prec(stack[stack.length - 1])) {
            result += stack.pop();
        }
        stack.push(s[i]);
    }

    while (stack.length !== 0) {
        result += stack.pop();
    }

    return result;
}

console.log(toPostfix("1+2+3"));
console.log(toPostfix("5+(6-2)*9+3^(7-1)"));
console.log(toPostfix("1^2^3"));
console.log(toPostfix("2+7*5"));
console.log(toPostfix("3*3/(7+1)"))
