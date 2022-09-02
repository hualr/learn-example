package _7_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, " doc: ", tagdoc)
	}
}

//其实就是java里的注解
func Test1(t *testing.T) {
	var re resume
	findTag(&re)
}
