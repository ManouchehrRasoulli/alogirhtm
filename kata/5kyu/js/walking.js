function collision(room) {
    var height = room.length;
    var width = room[0]?.length;
    if (!width || !height) return -1;
    var people = [];
    for (var y = 0; y < height; y++) {
        for (var x = 0; x < height; x++) {
            var char = room[y][x];
            if (char === "l") people.push({x: x, y: y, xmove: -1, ymove: 0});
            if (char === "r") people.push({x: x, y: y, xmove: 1, ymove: 0});
            if (char === "u") people.push({x: x, y: y, xmove: 0, ymove: -1});
            if (char === "d") people.push({x: x, y: y, xmove: 0, ymove: 1});
        }
    }
    var minDist = Infinity;
    for (var n = 0; n < people.length; n++) {
        var p1 = people[n];
        for (var n2 = 1; n2 < people.length; n2++) {
            var p2 = people[n2];
            if (p1.xmove === p2.xmove && p1.ymove === p2.ymove) continue;
            var dist = Math.abs(p1.x - p2.x) + Math.abs(p1.y - p2.y);
            if (dist >= minDist) continue;
            var newPos1 = {x: p1.x + p1.xmove * dist / 2, y: p1.y + p1.ymove * dist / 2};
            var newPos2 = {x: p2.x + p2.xmove * dist / 2, y: p2.y + p2.ymove * dist / 2};
            if (newPos1.x !== newPos2.x || newPos1.y !== newPos2.y) continue;
            minDist = dist;
        }
    }
    if (minDist === Infinity) return -1;
    return Math.ceil(minDist / 2);
}

const tests = [
    [
        [
            "-r--l".split(""),
            "-----".split(""),
            "-u-d-".split(""),
            "----l".split(""),
            "-----".split("")
        ], 1
    ],
    [
        [
            "-----".split(""),
            "--d-l".split(""),
            "-----".split(""),
            "rr--u".split(""),
            "--l--".split("")
        ], 2
    ],
    [
        [
            "-----".split(""),
            "---ur".split(""),
            "rd--l".split(""),
            "d-r--".split(""),
            "-r---".split("")
        ], 2
    ],
    [
        [
            "-u---".split(""),
            "---uu".split(""),
            "---r-".split(""),
            "--r--".split(""),
            "-d---".split("")
        ], -1
    ]
];

for (const [arr, exp] of tests) {
    console.log(collision(arr), exp);
}