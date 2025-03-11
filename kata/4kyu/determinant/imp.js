// determinant
// input.txt m :: matrix
function determinant(m) {
    if (m.length === 1) {
        return m[0][0];
    }

    if (m.length === 2) {
        return m[0][0] * m[1][1] - m[0][1]*m[1][0];
    }

    let res = 0;
    for (let col = 0; col < m.length; col++) {

        let sub = Array.from({length: m.length - 1}, ()=> new Array(m.length - 1));
        for (let i = 1; i < m.length; i++) {
            let subcol = 0;
            for (let j = 0; j < m.length;j++) {
                if (j === col) {
                    continue;
                }

                sub[i-1][subcol++] = m[i][j];
            }
        }

        let sign = (col % 2 === 0) ? 1:-1;
        res += sign * m[0][col] * determinant(sub);
    }

    return res;
}

let mat = [ [ 1, 0, 2, -1 ],
    [ 3, 0, 0, 5 ],
    [ 2, 1, 4, -3 ],
    [ 1, 0, 5, 0 ] ];

console.log(determinant(mat));
