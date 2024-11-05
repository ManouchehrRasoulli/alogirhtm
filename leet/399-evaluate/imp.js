/**
 * @param {string[][]} equations
 * @param {number[]} values
 * @param {string[][]} queries
 * @return {number[]}
 */
var calcEquation = function (equations, values, queries) {
    let graph = {};
    for (let i = 0; i < equations.length; i++) {
        let [a, b] = equations[i];
        let val = values[i];

        if (!graph[a]) graph[a] = {};
        if (!graph[b]) graph[b] = {};

        graph[a][b] = val;
        graph[b][a] = 1 / val;
    }

    function dfs(start, end, visited) {
        if (!(start in graph) || !(end in graph))
            return -1;

        if (start === end)
            return 1;

        visited.add(start);

        for (let ng in graph[start]) {
            if (visited.has(ng))
                continue;

            let resault = dfs(ng, end, visited);
            if (resault !== -1.0) {
                return resault * graph[start][ng];
            }
        }

        visited.delete(start);
        return -1.0;
    }

    let resault = [];
    for (let [a, b] of queries) {
        resault.push(dfs(a, b, new Set()));
    }

    return resault;
};

console.log(calcEquation([["a", "b"], ["b", "c"]], [2.0, 3.0], [["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"]]));
