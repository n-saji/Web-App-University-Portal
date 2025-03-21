package service

import (
	"CollegeAdministration/models"
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (ac *Service) ValidateLogin(email, password string) error {

	reg := regexp.MustCompile(`^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)
	ok := reg.MatchString(email)
	if !ok {
		return fmt.Errorf("wrong email format")
	} else if password == ":password" {
		return fmt.Errorf("password cant be empty ")
	} else if len(password) < 8 {
		return fmt.Errorf("password length insufficient")
	}
	return nil
}

func (s *Service) ValidateEmail(email string) error {
	ok, _ := regexp.MatchString("@gmail.com", email)
	if !ok || len(email) < 11 {
		return fmt.Errorf("wrong email format")
	}
	return nil
}
func (s *Service) ValidatePassword(password string) error {

	if password == "" {
		return fmt.Errorf("password cant be empty ")
	} else if len(password) < 8 {
		return fmt.Errorf("password length insufficient")
	}
	return nil
}
func (ac *Service) CheckEmailExist(email string) error {
	ok, err := ac.daos.CheckForEmail(email)
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf("email exists")
	}
	return nil
}
func (ac *Service) CheckCredentials(email, password string) error {

	hashed_password, err1 := ac.daos.FetchPasswordUsingEmailID(email)
	if err1 != nil {
		return err1
	}
	if hashed_password == "" {
		return fmt.Errorf("wrong email id")
	}
	err4 := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	if err4 != nil {
		return fmt.Errorf("wrong password")
	}

	return nil
}

func (ac *Service) GetTokenAfterLogging(account_id string) (uuid.UUID, error) {
	var token_table models.Token_generator
	account_uuid, err := uuid.Parse(account_id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error parsing uuid")
	}
	token_table.AccountId = account_uuid

	token := uuid.New()
	token_table.Token = token
	token_table.IsValid = true
	tn := time.Now()
	token_table.ValidFrom = tn.Unix()
	validtill := tn.Add(time.Hour * 24)
	token_table.ValidTill = validtill.Unix()

	err = ac.daos.InsertToken(token_table)
	if err != nil {
		return uuid.Nil, fmt.Errorf("token insertion failed")
	}
	return token_table.Token, nil
}

func (ac *Service) CheckTokenValidity(token uuid.UUID) (bool, error) {

	err1 := ac.CheckTokenExpiry(token)
	if err1 != nil {
		return false, err1
	}
	status, err := ac.daos.GetTokenStatus(token)
	if err != nil {
		return status, err
	}
	return status, nil
}

func (ac *Service) CheckTokenExpiry(token uuid.UUID) error {

	token_details, err := ac.daos.GetTokenStored(token)
	if err != nil {
		return err
	}
	time_now := time.Now()
	diff := token_details.ValidTill - time_now.Unix()
	if diff < 0 {
		err2 := ac.daos.SetTokenFalse(token)
		if err2 != nil {
			return err2
		}
		return fmt.Errorf("token expired! Generate new token")
	}

	return nil
}

func (s *Service) CheckTokenWithCookie(token string) error {

	token_id, err4 := uuid.Parse(token)
	if err4 != nil {
		return fmt.Errorf("error parsing uuid")
	}
	status, err1 := s.CheckTokenValidity(token_id)
	if err1 != nil {
		return err1
	}
	if !status {
		return fmt.Errorf("token expired")
	}
	return nil
}

func (s *Service) ValidateInstructorDetails(iid *models.InstructorDetailsDTO) (bool, error) {

	if iid.Department != "" {
		for _, each_letter := range iid.Department {
			if each_letter == '-' {
				return false, fmt.Errorf("wrong format! remove -")
			}
		}
	}

	return true, nil
}

func (s *Service) DisableToken(token string) error {

	parsedToken, err := uuid.Parse(token)
	if err != nil {
		return fmt.Errorf("unable to invalidate token, err- %s", err)
	}
	err = s.daos.DisableToken(parsedToken)
	if err != nil {
		return fmt.Errorf("db error while invalidating token, err - %s", err)
	}
	return nil
}

func (s *Service) GetAccountByToken(token string) (string, error) {

	parsedToken, err := uuid.Parse(token)
	if err != nil {
		return "", fmt.Errorf("unable to parse token, err- %s", err)
	}
	acc, err := s.daos.GetAccountByToken(parsedToken)
	if err != nil {
		return "", fmt.Errorf("db error while fetching account, err - %s", err)
	}
	return acc.AccountId.String(), nil
}
