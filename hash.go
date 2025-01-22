package CommonWealth

import (
	"github.com/go-crypt/crypt/algorithm"
	"github.com/go-crypt/crypt/algorithm/argon2"
)

// GenerateHashword takes a password and returns an encoded string w/ random salt
func GenerateHashword(password string) (string, error) {
	var (
		hasher *argon2.Hasher
		err    error
		digest algorithm.Digest
	)
	hasher, err = argon2.New(argon2.WithProfileRFC9106Recommended())
	if err != nil {
		return "", err
	}

	digest, err = hasher.Hash(password)
	if err != nil {
		return "", err
	}

	return digest.Encode(), nil
}
