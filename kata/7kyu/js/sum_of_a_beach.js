
function sumOfABeach(beach) {
    beach = beach.toLowerCase();
    let sand = beach.split('sand').length - 1;
    let water = beach.split('water').length - 1;
    let fish = beach.split('fish').length - 1;
    let sun = beach.split('sun').length - 1;
    
    return sand + water + fish + sun;
}

console.log(sumOfABeach("GolDeNSanDyWateRyBeaChSuNN"));