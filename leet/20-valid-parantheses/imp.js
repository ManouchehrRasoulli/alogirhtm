/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
    let v = new Array();

    for (let i = 0; i < s.length; i++) {
        switch (s[i]) {
            case "(":
            case "[":
            case "{":
                v.push(s[i]);
                break;
            default:
                {
                    if (v.length === 0)
                        return false;

                    let x = v.pop();
                    if (s[i] === "}" && x !== "{")
                        return false;

                    if (s[i] === ")" && x !== "(")
                        return false;

                    if (s[i] === "]" && x !== "[")
                        return false;
                }
        }
    }

    return v.length === 0;
};

console.log(isValid("()"));
console.log(isValid("()[]{}"));
console.log(isValid("(]"));

