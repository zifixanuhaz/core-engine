package core_engine

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// GenerateRandomString generates a random string of a specified length.
func GenerateRandomString(length int) (string, error) {
	const characterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const maxBigInt = big.NewInt(int64(len(characterSet)))

	var randomBigInt big.Int
	for i := 0; i < length; i++ {
		randomBigInt.Mul(&randomBigInt, maxBigInt)
		_, err := rand.Int(rand.Reader, maxBigInt)
		if err!= nil {
			return "", err
		}
		randomBigInt.Mod(&randomBigInt, maxBigInt)
	}

	var randomString string
	for i := 0; i < length; i++ {
		randomString += string(characterSet[randomBigInt.Int64()])
	}

	return randomString, nil
}

// GetTimestamp returns the current timestamp in seconds since the Unix epoch.
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// GenerateUUID generates a random UUID.
func GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err!= nil {
		return "", err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]),
		nil
}