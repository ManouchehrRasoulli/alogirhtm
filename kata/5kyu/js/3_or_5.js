function solution(number) {
    let index = 1;
    let x = 0;
    let y = 0;
    let sum = 0;
    while (x < number || y < number) {
        x = 3 * index;
        if (x < number) {
            if (x % 5 !== 0) {
                sum += x;
            }
        }
        y = 5 * index;
        if (y < number) {
            sum += y;
        }
        index++;
    }
    return sum;
}

console.log(solution(20));