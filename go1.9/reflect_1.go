package main

import (
	"fmt"
	"reflect"
)

const tagName = "default"

func main() {
	type User struct {
		BuildDate string `validate:"require"`
		Version   string `default:"1.9" validate:"require"`
		RPC struct {
			Host string `default:"0.0.0.0" validate:"require"`
		}
		LOG struct {
			Level string `default:"info" validate:"require"`
		}
		NET struct {
			Host string `default:"255.255.255.255" validate:"require"`
			Port string `default:"8000" validate:"require"`
		}
	}
	user := User{}

	// 获取tag中的内容
	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if v.Field(i).Type().Kind() != reflect.Struct {
			key := field.Name
			value := field.Tag.Get(tagName)
			fmt.Printf("key is: %v,value is: %v\n", key, value)
		} else {
			structField := v.Field(i).Type()
			for j := 0; j < structField.NumField(); j++ {
				key := structField.Field(j).Name
				values := structField.Field(j).Tag.Get(tagName)
				fmt.Printf("key is: %v,value is: %v\n", key, values)
			}
			continue
		}
	}
}
