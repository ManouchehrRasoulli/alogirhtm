function ipToInt32(ip) {
    let parts = ip.split(".");

    let v = (Number(parts[0]) * Math.pow(256, 3)) |
        (Number(parts[1]) * Math.pow(256, 2)) |
        (Number(parts[2]) * Math.pow(256, 1)) |
        (Number(parts[3] * Math.pow(256, 0)));

    return v >>> 0;
}

console.log(ipToInt32("128.32.10.1")) // 2149583361
