// code5.1
// Frog問題を動的計画法で解く

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var INF float64 = math.Pow(2, 60)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	h := []float64{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < N; i++ {
		scanner.Scan()
		in, _ := strconv.Atoi(scanner.Text())
		h = append(h, float64(in))
	}

	// スライス全体を無限大を表す値で初期化
	dp := make([]float64, N)
	for i := 0; i < N; i++ {
		dp = append(dp, INF)
	}

	// 初期条件
	dp[0] = 0

	for i := 1; i < N; i++ {
		if i == 1 {
			dp[i] = math.Abs(h[i] - h[i-1])
		} else {
			dp[i] = math.Min(dp[i-1]+math.Abs(h[i]-h[i-1]), dp[i-2]+math.Abs(h[i]-h[i-2]))
		}
	}

	fmt.Println(dp[N-1])
}
