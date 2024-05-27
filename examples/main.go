package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	input := bufio.NewScanner((os.Stdin))
	counter := 0
	for input.Scan() {
		counts[input.Text()]++
		counter++
		if counter == 5 {
			break
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
		}
	}
}
