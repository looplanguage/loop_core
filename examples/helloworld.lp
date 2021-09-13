var loop = fun(arr, cb) {
    cb(arr)
}

loop([1, 2, 3], fun(a) {
    println("hey")
})