
function _abbrevName(name){
    return name.split(' ').map(i => i[0].toUpperCase()).join('.')
}

function abbrevName(name){
    name = name.toUpperCase().split(" ");
    return name[0][0] + '.' + name[1][0];
}

console.log(_abbrevName("manouchehr rasouli"))