package encryption

import (
	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/chacha20poly1305"
)

func Encrypt(data, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, chacha20.NonceSizeX)
	return aead.Seal(nil, nonce, data, nil), nil
}
