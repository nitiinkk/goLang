package main
import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //declared two variables s, sep of type string
	for i:=1; i<len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println("the resultant string is", s);
}

// go run 2.osArgs.go hello world
//output : the resultant string is hello world