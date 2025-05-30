package auth

import "testing"

func TestCreateToken(t *testing.T) {
	userID := "12345"
	secretStr := "secret"

	secret := []byte(secretStr)
	token, err := GenerateJWT(userID, secret, "user")
	if err != nil {
		t.Fatal(err)
	}

	if token == "" {
		t.Error("Token is empty")
	}
}