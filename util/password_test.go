package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	passwd := RandomString(6)
	hashedPasswd, err := HashPassword(passwd)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPasswd)

	err = CheckPasswordHash(passwd, hashedPasswd)
	require.NoError(t, err)

	wrongPasswd := RandomString(7)
	err = CheckPasswordHash(wrongPasswd, hashedPasswd)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	
}
