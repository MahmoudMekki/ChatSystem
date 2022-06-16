package utils

import (
	"crypto/rand"
	"github.com/MahmoudMekki/ChatSystem/config"
	"math/big"
	"strconv"
)

func GenerateRandomString() (string, error) {
	sizeStr := config.GetEnvVar("TOKEN_SIZE")
	size, _ := strconv.Atoi(sizeStr)
	tok := make([]byte, size)
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < size; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		tok[i] = letters[num.Int64()]
	}
	return string(tok), nil
}
