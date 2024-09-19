var cutLog = function (p, n) {
    let v = [];
    v.push(0);
    for (j = 1; j <= n; j++) {
        let q = -100000;
        for (i = 1; i <= j; i++) {
            q = Math.max(q, p[i]+v[j-i]);
        }
        v[j]=q;
    }
    return v[n];
}
