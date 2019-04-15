package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("parse a simple ineteger")
	var one MyInt = 1
	walk(one, 0)

	fmt.Println("parse a simple struct")
	two := struct{ Name string }{"foo"}
	walk(two, 0)

	fmt.Println("parse a recursive struct")
	p := &Person{
		Name:    &Name{"stephen", "lebron", "james"},
		Address: &Address{"bay area", "california"},
	}
	walk(p, 0)

}

type MyInt int

type Person struct {
	Name    *Name
	Address *Address
}

type Name struct {
	Title string
	First string
	Last  string
}

type Address struct {
	Area  string
	State string
}

func walk(u interface{}, depth int) {
	val := reflect.Indirect(reflect.ValueOf(u))
	t := val.Type()

	tabs := strings.Repeat("\t", depth+1)

	fmt.Printf("%sthe type of %s is %q \n", tabs, val.Kind(), t)

	if val.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldVal := reflect.Indirect(val.Field(i))

			tabs := strings.Repeat("\t", depth+2)
			fmt.Printf("%sFIeld %q is type %q (%s) \n", tabs, field.Name, field.Type, fieldVal.Kind())

			if fieldVal.Kind() == reflect.Struct {
				walk(fieldVal.Interface(), depth+1)
			}
		}
	}

}
