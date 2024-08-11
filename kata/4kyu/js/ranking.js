class User {
    rank; // -8 --> +8
    progress; // 0 -> 100 each time get over 100 rank++

    constructor() {
        console.log("new user created !");
        this.rank = -8;
        this.progress = 0;
    };

    _dif_(rank) {
        let count = 0;
        if (rank < this.rank) {
            for (let i = rank; i < this.rank; i++) {
                if (i === 0) {
                    continue;
                }
                count++;
            }
            return count;
        }

        for (let i = this.rank; i < rank; i++) {
            if (i === 0) {
                continue;
            }
            count++;
        }
        return count;
    };

    _add_progress_(x) {
        console.log("add progress --> rank:", this.rank, "point:", this.progress, "x:", x);
        this.progress += x;
        while (this.progress >= 100) {
            this.progress -= 100;
            this.rank++;
            if (this.rank === 0) {
                this.rank = 1;
            }
            console.log(this.progress, this.rank);
        }
        if (this.rank >= 8) {
            this.rank = 8;
            this.progress = 0;
        }
        console.log("after adding progress --> rank:", this.rank, "point:", this.progress);
    };

    incProgress(rank) {
        if (rank === 0 || rank < -8 || rank > 8) {
            throw new Error("invalid rank");
        }
        console.log("solve rank:", rank, "question with rank:", this.rank);
        if (rank <= this.rank) {
            let dif = this._dif_(rank);
            if (dif > 2) {
                return;
            }

            if (dif === 1) {
                this._add_progress_(1);
                return;
            }

            if (dif === 0) {
                this._add_progress_(3);
                return;
            }
        }
        let dif = this._dif_(rank);
        console.log("diff is: ", dif);
        let x = (dif * dif) * 10;
        this._add_progress_(x);
    };
}

let user = new User();
console.log(user.rank);
console.log(user.progress);
user.incProgress(1);
console.log(user.rank);
console.log(user.progress);
