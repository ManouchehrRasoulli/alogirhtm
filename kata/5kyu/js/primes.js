function primeFactors(n) {

    let items = [];

    while (n % 2 === 0) {
        items.push(2);
        n = Math.floor(n / 2);
    }

    for (let i = 3;
         i <= Math.floor(Math.sqrt(n));
         i = i + 2) {

        while (n % i === 0) {
            items.push(i);
            n = Math.floor(n / i);
        }
    }

    // greater than 2
    if (n > 2)
        items.push(n);

    let str = "";

    console.log(items);

    for (let i = 0; i < items.length; i++) {

        let x = "(" + items[i];
        let count = 0;

        let j = 0;
        for (j = i; items[j] === items[i]; j++) {
            if (items[j] === items[i]) {
                count++;
            }
        }
        i = j - 1;

        if (count > 1) {
            x += "**" + count;
        }
        x += ")";

        str += x;
    }

    return str;
}

console.log(primeFactors(7775460));
