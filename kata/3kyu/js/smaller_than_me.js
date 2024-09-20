function merge(v, ans, l, mid, h) {

    let t = []; // temporary array for merging both halves
    let i = l;
    let j = mid + 1;

    while (i < mid + 1 && j <= h) {

        // v[i].first is greater than all
        // the elements from j till h.
        if (j < v.length && v[i][0] > v[j][0]) {
            ans[v[i][1]] += (h - j + 1);
            t.push(v[i]);
            i++;
        } else {
            t.push(v[j]);
            j++;
        }
    }

    // if any elements left in left array
    while (i <= mid) {
        t.push(v[i]);
        i++;
    }

    // if any elements left in right array
    while (j <= h) {
        t.push(v[j]);
        j++;
    }

    // putting elements back in main array in
    // descending order
    for (let k = 0, i = l; i <= h; i++, k++)
        v[i] = t[k];
}

function mergesort(v, ans, i, j) {
    if (i < j) {
        let mid = Math.round((i + j) / 2) - 1;

        // calling mergesort for left half
        mergesort(v, ans, i, mid);

        // calling mergesort for right half
        mergesort(v, ans, mid + 1, j);

        // merging both halves and generating answer
        merge(v, ans, i, mid, j);
    }
}

function smaller(arr) {

    let v = [];
    let n = arr.length;

    // inserting elements and corresponding index
    // as pair
    for (let i = 0; i < n; i++) {
        let x = [arr[i], i];
        v.push(x);
    }

    // answer array for keeping count
    // initialized by 0,
    let ans = new Array(n);
    for (let i = 0; i < n; i++)
        ans[i] = 0;

    // calling mergesort
    mergesort(v, ans, 0, n - 1);

    return ans;
}

console.log(smaller([1, 2, 3]));
