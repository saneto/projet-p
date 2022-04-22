package security

import "golang.org/x/crypto/bcrypt"

// Hash : simple password hashing method
func Hash(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword : This method compare the password with the hash
func VerifyPassword(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}
