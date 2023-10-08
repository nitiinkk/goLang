package main
import (
	"fmt"
	"os"
	"net/http"
	"io"
	"io/ioutil"
	"time"
)

func main() {
	start:=time.Now()
	ch:=make(chan string) //makes a channel in Go
	for _, url:= range os.Args[1:] {
		go fetch(url, ch) //starts a go routine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch (main receives from channel ch)
	}

	fmt.Printf("%.2f seconds elapsed \n", time.Since(start).Seconds())

}

func fetch (url string, ch chan<-string) { //ch is of type channel which only receives string
	start:=time.Now()
	resp, err :=http.Get(url)
	if err!=nil {
		ch<-fmt.Sprint(err)
		return
	}

	nbytes, err:=io.Copy(ioutil.Discard, resp.Body) //Copy returns the response byte
	resp.Body.Close()
	if err!=nil {
		ch<-fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs:=time.Since(start).Seconds()
	ch<-fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

//input go run 6.fetchAll.go https://golang.org http://gopl.io https://godoc.org
/* output :
2.13   61870 https://golang.org
2.71   31630 https://godoc.org
4.01    4154 http://gopl.io
4.01 seconds elapsed 
*/