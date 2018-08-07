package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("handler error..")
	}

	//note:在使用完成之后，切记要将程序的主体关闭。
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

//HTTP GET请求的实现
