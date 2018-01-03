package main

import (
	"github.com/tidwall/gjson"
	"fmt"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json, "name.last")
	fmt.Println("values:", value.String())
}

