function isToday(date) {
    var today = new Date();
    console.log(date, today);

    if (today.getUTCFullYear() !== date.getUTCFullYear() ||
        today.getUTCMonth() !== date.getUTCMonth() ||
        today.getDay() !== date.getDay() ||
        today.getUTCDate() !== date.getUTCDate()) {

        return false;
    }

    return true;
}

const today = new Date();
const yesterday = new Date();
yesterday.setDate(today.getDate() - 1);

const dayAfter = new Date();
dayAfter.setDate(today.getDate() + 7)

console.log(isToday(yesterday));
console.log(isToday(dayAfter));
console.log(isToday(today));
