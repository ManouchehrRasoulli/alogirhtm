function dashatize(num) {
    let result = "";
    String(num).split('').forEach((value, index, arr) => {
        let n = Number(value);
        if (n % 2 === 1) {
            if (index !== 0 && index !== arr.length - 1) {
                result = result + "-" + n + "-";
            } else if (index === 0) {
                result = n + "-";
            } else {
                result = result + "-" + n;
            }
            return;
        }
        result = result + n;
    });
    result = result.replaceAll("--", "-");
    return result;
}

console.log(dashatize(5311));
