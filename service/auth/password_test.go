package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "pwd"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}

	if hash == "" {
		t.Error("Hash password is empty")
	}

	if hash == password {
		t.Error("Hash password is equal to password")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "pwd"
	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}

	if !CheckPassword(password, hashed) {
		t.Error("Password is not equal to hashed password")
	}
}