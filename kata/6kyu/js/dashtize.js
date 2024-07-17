function dashatize(num) {
    if (num < 0) {
        num = num * -1;
    }
    let result = "";
    String(num).split('').forEach((value, index, arr) => {
        let n = Number(value);
        if (n % 2 === 1) {
            if (index !== 0 && index !== arr.length - 1) {
                result = result + "-" + n + "-";
            } else if (index === 0 && arr.length === 1) {
                result = n + "";
            } else if (index === 0) {
                result = n + "-";
            } else {
                result = result + "-" + n;
            }
            return;
        }
        result = result + "" + n;
    });
    return result.replaceAll("--", "-");
}

console.log(dashatize(-1));
console.log(dashatize(222));
console.log(dashatize(274));
console.log(dashatize(5311));
