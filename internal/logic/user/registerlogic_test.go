package user

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestRegisterLogic_Register(t *testing.T) {
	password := "111"
	passwordOk := "$2a$10$8Apn7CKJMwxNQ.3/wOzXcu3iCZM51RKcmvNoeGaKz2YV/2lIoQIBu"

	err := bcrypt.CompareHashAndPassword([]byte(passwordOk), []byte(password))
	if err != nil {
		t.Error(err)
	}
}
