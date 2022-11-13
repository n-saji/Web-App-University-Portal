package service

import (
	"fmt"
	"regexp"
)

func (ac *Service) ValidateLogin(email, password string) error {

	ok, _ := regexp.MatchString("@gmail.com", email)
	if !ok {
		return fmt.Errorf("wrong email format")
	} else if password == ":password" {
		return fmt.Errorf("password cant be empty ")
	} else if len(password) < 8 {
		return fmt.Errorf("password length insufficient")
	}
	return nil
}
