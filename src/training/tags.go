package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `common_tag:"name_common_tag_val" extra_tag:"extra_tag_val"`
	Email string `common_tag:"email_common_tag_val, omitempty"`
}

func main() {
	user := User{Name: "Christopher", Email: "christopher1992@gmail.com"}
	t := reflect.TypeOf(user)

	for _, fieldName := range []string{"Name", "Email"} {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}
		fmt.Printf("\nField: User.%s \n", fieldName)
		fmt.Printf("\tWhole tag value: %q \n", field.Tag)
		fmt.Printf("\tValue of common_tag: %q \n", field.Tag.Get("common_tag"))
		fmt.Printf("\tValue of extra_tag: %q \n", field.Tag.Get("extra_tag"))
	}
}
