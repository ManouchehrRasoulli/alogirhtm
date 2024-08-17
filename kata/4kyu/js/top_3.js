function topThreeWords(text) {
    let words = {}
    text.toLowerCase().replace(/([A-Za-z][A-Za-z']*)/g, match => {
        let c = words[match] || 0
        words[match] = ++c
    })
    return Object
        .keys(words)
        .sort(function(a,b){return words[b]-words[a]})
        .slice(0,3)
}

console.log(topThreeWords("a a a  b  c c  d d d d  e e e e e"), ['e','d','a']);