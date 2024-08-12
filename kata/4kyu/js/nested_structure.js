Array.prototype.sameStructureAs = function (other) {
    if (this.length !== other.length) {
        return false;
    }

    for (let i = 0; i < this.length; i++) {
        if (Array.isArray(this[i])) {
            if (!this[i].sameStructureAs(other[i])) {
                return false;
            }
        } else if (Array.isArray(other[i])) {
            return false;
        }
    }

    return true;
};

console.log([1, [1, 1]].sameStructureAs([2, [2, 2]]));
