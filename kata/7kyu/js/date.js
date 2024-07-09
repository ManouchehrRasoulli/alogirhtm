function findYear(month, dayOfWeek) {

    for (let year = 2014; year <= 2050; ++year) {
        let firstDayOfMonth = new Date(year, month, 1);

        if (firstDayOfMonth.getDay() === dayOfWeek) {
            return year;
        }
    }

    return 0;
}

console.log(findYear(11, 2));
console.log(findYear(4, 55));