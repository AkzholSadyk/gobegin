package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func minCost(n, k int, a []int) int {
	sort.Ints(a) // Сортируем массив

	// Берём k/2 самых маленьких чисел
	b := append(a[:k/2], 0) // Добавляем 0 в конец

	// Вычисляем стоимость b
	for i := 0; i < len(b); i++ {
		if b[i] != i+1 {
			return i + 1
		}
	}
	return len(b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n, k int
		fmt.Fscan(reader, &n, &k)

		a := make([]int, n)
		for i := range a {
			fmt.Fscan(reader, &a[i])
		}

		fmt.Println(minCost(n, k, a))
	}
}
