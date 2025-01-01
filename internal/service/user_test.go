package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassswordEncrypt(t *testing.T) {
	password := []byte("123456")
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.NoError(t, err)
	println(string(encrypted))
	err2 := bcrypt.CompareHashAndPassword(encrypted, []byte("123456"))
	fmt.Println("err", err2)
	//assert.NotNil(t, err2)
	assert.NoError(t, err2)
}
