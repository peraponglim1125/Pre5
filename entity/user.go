package entity

import "github.com/asaskevich/govalidator"

type User struct {
	Username string `valid:"required~Username is required"` 
    Password string  `valid:"stringlength(8|1999)~Password must be at least 8 characters"` 
    Email    string   `valid:"required~Email is required, email~Invalid email format"` 
    Age      int  `valid:"range(13|100)~Age must be at least 13"` 
}

func ValidationUser(user *User) (bool, error){
	return govalidator.ValidateStruct(user)

}