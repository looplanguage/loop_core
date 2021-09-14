import "./examples/packages/math.lp" as math
import "./examples/packages/std.lp" as _
import "./examples/packages/number.lp" as number

for([1, 2, 3], fun(key, value) {
  println(key, ": ", value)
})

math["double"](_Test3)