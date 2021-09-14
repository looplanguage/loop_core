var for = fun(arr, cb) {
    var index = 0

    while(index < len(arr) - 1) {
        cb(index, arr[index])
        var index = index + 1
    }
}

export {
  "for": for
}