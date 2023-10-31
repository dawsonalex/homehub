package id

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

const (
	idPrefixLength = 5
	idLength       = 21
)

var (
	ErrBadIdLength    = errors.New("id incorrect length")
	ErrPrefixNotFound = errors.New("id prefix doesn't exist")
)

type Id string

func (id Id) Prefix() string {
	return string(id[:idPrefixLength])
}

func NewFromSequence(prefix string, seq int64) (Id, error) {
	prefix = padPrefix(prefix)
	typeKey, typeExists := registry.typeMap[prefix]
	if !typeExists {
		return "", ErrPrefixNotFound
	}

	aesBlock, err := aes.NewCipher([]byte(typeKey))
	if err != nil {
		return "", err
	}

	toEncrypt := make([]byte, 8)
	binary.PutVarint(toEncrypt, seq)

	// The IV needs to be unique, but not secure. Therefore, it's common to
	//include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize)
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// encrypt the sequence bytes
	cipherBytes := make([]byte, len(toEncrypt))
	stream := cipher.NewCTR(aesBlock, iv)
	stream.XORKeyStream(cipherBytes, toEncrypt)

	// return the ID as the prefix concat with the hex encoded cipher.
	return Id(fmt.Sprintf("%s%s", prefix, hex.EncodeToString(cipherBytes))), nil
}

// TODO: Is ToSequence required? If so might need to store the iv at encryption.
//func ToSequence(id Id) (int64, error) {
//	decodedCipher, err := hex.DecodeString(string(id))
//	if err != nil {
//		return 0, err
//	}
//
//	cipherKey, exists := registry.get(id.Prefix())
//	if !exists {
//		return 0, ErrPrefixNotFound
//	}
//
//
//}

func IsValid(id Id) (bool, error) {
	if len(id) != idLength {
		return false, ErrBadIdLength
	}

	if _, isRegistered := registry.get(id.Prefix()); !isRegistered {
		return false, ErrPrefixNotFound
	}

	return true, nil
}

func padPrefix(prefix string) string {
	return fmt.Sprintf("%s00000", prefix)[:idPrefixLength]
}
