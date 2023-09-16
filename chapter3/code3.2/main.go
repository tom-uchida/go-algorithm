package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	foundId := -1
	for i := 0; i < N; i++ {
		if a[i] == v {
			foundId = i
			break
		}
	}

	fmt.Println(foundId)
}
