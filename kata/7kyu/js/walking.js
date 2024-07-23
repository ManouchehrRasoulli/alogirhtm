function contact(hallway) {
    let loc = [];
    for (let i = 0; i < hallway.length; i++) {
        if (hallway[i] === '-') {
            continue;
        }

        (hallway[i] === '<') ? loc.push(-(i + 1)) : loc.push(i + 1);
    }

    let bestDif = hallway.length + 1;

    for (let i = 0; i < loc.length; i++) {
        if (loc[i] < 0) {
            continue;
        }

        for (let j = i + 1; j < loc.length; j++) {
            if (loc[j] > 0) {
                continue;
            }

            let dif = (-loc[j] - loc[i]) / 2;
            if (dif < bestDif) {
                bestDif = dif;
            }
            break;
        }
    }

    return (bestDif > hallway.length) ? -1 : Math.ceil(bestDif);
}

console.log(contact("---->---<"), 2);
console.log(contact("----<-->---"), -1);
console.log(contact("----><-----"), 1);
console.log(contact(">>>>--<<<<<<<<<<<<<<<<<<<<"), 2);
console.log(contact(">---------------<--------------------------<-------->---------<------->----------<----<---->>----------<------->>>---------------<<------>"), 5);
console.log(contact("<><>"), 1);
console.log(contact("<<<<<<<<>>>>>>>>"), -1);