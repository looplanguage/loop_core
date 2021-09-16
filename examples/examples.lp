import "./examples/packages/math.lp" as math // Imports local package as math
import "https://github.com/looplanguage/std" as _ // Imports remote package as root (latest version) (will cache)
import "./examples/packages/number.lp" as number // Imports local package as number

for([1, 2, 3], fun(key, value) {
  println(key, ": ", value)
})

math["double"](100) // Still working on support for math.double instead of []
