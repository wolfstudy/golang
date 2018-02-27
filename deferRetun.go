package main

import "fmt"

func Nuhsqwdq() (ret bool) {
	ret = false
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			ret = true
		}
	}()
	panic("helllo")
	return ret
}

func main() {

}
