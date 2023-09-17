// code4.6
// フィボナッチ数列を求める再帰関数の再帰呼び出しの様子

package main

import "fmt"

func main() {
	fibo(6)
}

func fibo(N int) int {
	fmt.Printf("fibo(%d) を呼び出しました\n", N)

	// ベースケース
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}

	result := fibo(N-1) + fibo(N-2)
	fmt.Printf("%d 項目 = %d\n", N, result)

	return result
}
