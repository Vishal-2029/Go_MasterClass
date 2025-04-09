package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string `validate:"required,min=5,max=20"`
	Email    string `validate:"required,email"`
}

func validateStruct(data interface{}) []error {
	errs := []error{}
	v := reflect.ValueOf(data)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tag := t.Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}
		if validateField(v.Field(i).Interface(), tag) {
			continue
		}
		errs = append(errs, fmt.Errorf("%s validation failed", t.Field(i).Name))
	}
	return errs
}

func validateField(val interface{}, tag string) bool {
	
	return true
}

func main() {
	user := User{
		Username: "Vishal",
		Email:    "vishal@gmail.com",
	}
	errs := validateStruct(user)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Validation passed")
	}
}