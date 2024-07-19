function alphanumeric(string) {
    if (string.length === 0) {
        return false;
    }
    return !/[^a-zA-Z0-9]/.test(string);
}

console.log(alphanumeric("Mazinkaiser"), true);
console.log(alphanumeric("hello world_"), false);
console.log(alphanumeric("PassW0rd"), true);
console.log(alphanumeric("     "), false);
console.log(alphanumeric(""), false);