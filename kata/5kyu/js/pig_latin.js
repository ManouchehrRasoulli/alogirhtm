function pigIt(str) {
    console.log(str);
    let pigStr = "";
    str.split(" ").forEach(function (param) {
        if (param.length === 1 && !/[a-zA-Z]/.test(param)) {

            pigStr += param + " ";
            return;
        }
        pigStr += param.slice(1) + param[0] + "ay ";
    });
    return pigStr.trimEnd();
}

console.log(pigIt('Pig latin is cool'));
console.log(pigIt('O tempora o mores !'));
