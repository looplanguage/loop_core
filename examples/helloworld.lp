var for = fun(arr, cb) {
    var index = 0

    while(index < len(arr) - 1) {
        cb(arr[index], index)
        var index = index + 1
    }

    return index
}

for(["omg", "it", "works", "!", 0], fun(key, value) {
})