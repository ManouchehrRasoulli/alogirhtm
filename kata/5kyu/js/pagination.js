class PaginationHelper {
    constructor(collection, itemsPerPage) {
        this._count_ = collection.length;
        this._item_per_page_ = itemsPerPage;
        console.log(this._count_, this._item_per_page_);
    }

    itemCount() {
        return this._count_;
    }

    pageCount() {
        return Math.ceil(this._count_ / this._item_per_page_);
    }

    pageItemCount(pageIndex) {
        if (this._count_ === 0) {
            return -1;
        }
        let startIndex = pageIndex * this._item_per_page_;
        if (startIndex >= this._count_ || startIndex < 0) {
            return -1;
        }
        let endIndex = (pageIndex + 1) * this._item_per_page_ - 1;
        return (endIndex < this._count_) ? endIndex - startIndex + 1 : this._count_ - startIndex;
    }

    pageIndex(itemIndex) {
        if (this._count_ === 0) {
            return -1;
        }
        return (itemIndex >= 0 && itemIndex < this._count_) ? Math.floor(itemIndex / this._item_per_page_) : -1;
    }
}

const collection = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24];
const helper = new PaginationHelper(collection, 10);

console.log(helper.itemCount());
console.log(helper.pageCount());
console.log(helper.pageItemCount(2));
console.log(helper.pageIndex(9));
