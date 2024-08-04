var alp = "ABCDEFGHIJKLMNOPQRSTUVWXYZ ?!@#&()|<>.:=-+*/0123456789";

function rotorChange(line, changes) {
    let rotorSum = 0;
    for (let i = 0; i < line.length; i++) {
        rotorSum += changes[i];
        changes[i] = alp.indexOf(line[i]) + rotorSum;
    }

    return changes;
}

function route(line, changes) {
    let l = "";
    let ch = rotorChange(line, changes);
    for (let i = 0; i < line.length; i++) {
        l += alp[changes[i] % alp.length]
    }
    return l;
}

var flapDisplay = function (lines, rotors) {
    let after = [];
    for (let i = 0; i < lines.length; i++) {
        after[i] = route(lines[i], rotors[i])
    }
    return after;
}

console.log(flapDisplay(["HELLO ", "CODE"], [[15, 49, 50, 48, 43, 13], [20, 20, 28, 0]])); // [WORLD!, WARS]
