let j = {name: 'Time', age: 22}
let jj = JSON.stringify(j)
let ji = JSON.parse(jj, function (k, v) {
    console.log(k, v)
    return v+"22"
})

console.log(ji)