import "./examples/packages/math.lp" as math
import "https://github.com/looplanguage/std" as _
import "./examples/packages/number.lp" as number

for([1, 2, 3], fun(key, value) {
  println(key, ": ", value)
})

double(100)
