function solution(text, markers) {
    if (markers.length === 0) {
        return text.trimEnd();
    }

    let trimmedText = "";
    for (let i = 0; i < text.length; i++) {
        if (markers.indexOf(text[i]) === -1) {
            trimmedText += text[i];
            continue;
        }

        while (text[i] !== '\n' && i < text.length) {
            i++;
        }
        if (i < text.length) {
            trimmedText += '\n';
        }
    }

    let i = trimmedText.length-1;
    while (trimmedText[i] === ' ') {
        i--;
    }

    return trimmedText.substring(0, i+1).replaceAll(' \n', '\n');
}

console.log(solution('#aa bb cc', ['#']) === '');
console.log(solution('aa ! bb\ncc # dd', ['#', '!']) === 'aa\ncc');
console.log(solution('aa bb\n#cc dd', ['#']) === 'aa bb\n');