// code4.2
// 1からNまでの総和を計算する再帰関数

package main

import (
	"fmt"
)

func main() {
	f(5)
}

func f(N int) int {
	fmt.Printf("func(%d) を呼び出しました\n", N)

	if N == 0 {
		return 0
	}

	result := N + f(N-1)
	fmt.Printf("%d までの和 = %d\n", N, result)

	return result
}
