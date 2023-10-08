package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts:=make(map[string]int) //creates a Map with key as string and value as int
	input:=bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++;
	}

	// fmt.Println(counts)

	for line, n:=range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//input : go run 4.print\ input\ text.go hello 1 hello 1
//output : 2 hello 1
//map[hello1:2 we are:1]