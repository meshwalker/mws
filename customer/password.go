package customer

import (
	"golang.org/x/crypto/scrypt"
	log "github.com/Sirupsen/logrus"
	"crypto/rand"
	"bytes"
)

func GenSaltedPassword(password string) (key, salt []byte){
	var err error

	if salt, err = GenSalt(64); err != nil {
		log.Fatal("Unable to salt password --> ", err)
	}

	if key, err = scrypt.Key([]byte(password), salt, 16384, 8, 1, 32); err != nil {
		log.Fatal("Unable to salt password --> ", err)
	}

	return key, salt
}


func ComparePasswords(storedPassword []byte, password string, salt []byte) bool {
	comp, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
	if err != nil {
		log.Fatal("Unable to compare passwords --> ", err)
	}

	if bytes.Equal(storedPassword, comp) {
		return true
	}

	return false
}


func GenSalt(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}