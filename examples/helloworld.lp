var for = fun(arr, cb) {
    var index = 0

    while(index < len(arr) - 1) {
        cb(index, arr[index])
        var index = index + 1
    }
}

for(["omg", "it", "works", "!", 0], fun(key, value) {
    println(key, ":1 ", value)
})