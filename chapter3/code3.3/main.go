package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const INF = 20000000

func main() {
	var N, v int
	fmt.Scanf("%d %d", &N, &v)
	a := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < N; i++ {
		scanner.Scan()
		in, _ := strconv.Atoi(scanner.Text())
		a = append(a, in)
	}

	minValue := INF
	for i := 0; i < N; i++ {
		if a[i] < minValue {
			minValue = a[i]
		}
	}

	fmt.Println(minValue)
}
