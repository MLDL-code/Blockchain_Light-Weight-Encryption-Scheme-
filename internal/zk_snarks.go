package zkpee

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func GenerateProof(data []byte) []byte {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	proof, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, 0, data)
	return proof
}

func VerifyProof(proof, data []byte, pubKey *rsa.PublicKey) bool {
	err := rsa.VerifyPKCS1v15(pubKey, 0, data, proof)
	return err == nil
}
 
