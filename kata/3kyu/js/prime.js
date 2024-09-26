function isPrime(n) {
    if (n <= 1)
        return false;
    if (n === 2 || n === 3)
        return true;
    if (n % 2 === 0 || n % 3 === 0)
        return false;
    for (let i = 5; i <= Math.sqrt(n); i = i + 6)
        if (n % i === 0 || n % (i + 2) === 0)
            return false;
    return true;
}

function nextPrime(current) {
    let x = current + 1;
    while (true) {
        if (isPrime(x)) {
            return x;
        } else {
            x++;
        }
    }
}

class Primes {
    static* stream() {
        let current = 2;
        while (true) {
            yield current;
            current = nextPrime(current);
        }
    }
}


const stream = Primes.stream()
for (let i = 0; i < 1_000_000; i++) {
    stream.next();
}

console.log(stream.next());
