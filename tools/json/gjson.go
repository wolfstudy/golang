package main

import "github.com/gjson"

func main() {
	json := `{
	          "name":
					{
                           "first":"Janet",
                           "last":"Prichard"
					},
              "age":47}`
	value := gjson.Get(json, "name.last")
	println(value.String())

	json1 := `
					{
	  "programmers": [
		{
		  "firstName": "Janet",
		  "lastName": "McLaughlin",
		}, {
		  "firstName": "Elliotte",
		  "lastName": "Hunter",
		}, {
		  "firstName": "Jason",
		  "lastName": "Harold",
		}
	  ]
	}
`
	//Suppose you want all the last names from the following json:
	result := gjson.Get(json1, "programmers.#.lastName")
	for _, name := range result.Array() {
		println(name.String())
	}
}
