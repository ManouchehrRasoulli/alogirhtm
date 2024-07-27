function generateHashtag(str) {
    let sp = str.split(" ");
    let hashTag = "";
    sp.filter(function (param) {
        if (param === "") {
            return;
        }
        hashTag = hashTag + param[0].toUpperCase() + param.slice(1);
    })
    if (hashTag.length === 0 || hashTag.length > 139) {
        return false;
    }

    return "#" + hashTag;
}

console.log(generateHashtag("some test string"));
console.log(generateHashtag("   some   s x"));
console.log(generateHashtag("Codewars Is Nice"));
console.log(generateHashtag("code" + " ".repeat(140) + "wars"));
