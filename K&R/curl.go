package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)


// ./curl http://www.baidu.com(可获取百度的body，类似于curl的作用)
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reding:%s %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}
