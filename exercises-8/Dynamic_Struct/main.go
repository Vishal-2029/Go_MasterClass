package main

import v "github.com/go-playground/validator/v10"

type User struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func main() {
	validator := v.New()

	user := User{
		Name : "Vishal",
		Email: "",
	}

	err := validator.Struct(user)
	if err !=nil{
		for _, err := range err.(v.ValidationErrors){
			println("Field:", err.Field())
			println("Tag:", err.Tag())
		}
	}else{
		println("Validation passed")
	}
}