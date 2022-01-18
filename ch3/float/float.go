package main

import (
	"fmt"
	"math"
)

func main() {
	{
		var f float32 = 16777216   // 1 << 24
		fmt.Println(f == f+1)      // true !!
		var f64 float64 = 16777216 // 1 << 24
		fmt.Println(f64 == f64+1)  // false OK!!
	}

	// %gヴァーブは適切な精度で表現
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d ex = %8.3f %%g %g \n", x, math.Exp(float64(x)), math.Exp(float64(x)))
	}

	// NaN (not a number)
	// Inf (Infinity)
	{
		var z float64
		fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN
	}

	// NaNとの比較は常にfalse
	{
		nan := math.NaN()
		fmt.Println(nan == nan, nan < nan, nan > nan) // false false false
	}

}
