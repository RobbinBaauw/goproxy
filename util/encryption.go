package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
)

var DERPublicKey = getDERKey()
var PrivateKey *rsa.PrivateKey

var VerifyToken = getVerifyToken()

var reader = rand.Reader

func getVerifyToken() []byte {
	verifyTokenBytes := make([]byte, 4)
	_, _ = rand.Read(verifyTokenBytes)
	return verifyTokenBytes
}

func getDERKey() []byte {
	key, err := rsa.GenerateKey(reader, 1024)
	checkError(err)

	PrivateKey = key

	asn1Bytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	checkError(err)

	return asn1Bytes
}

func checkError(err error) {
	if err != nil {
		log.Panic("Fatal error ", err.Error())
	}
}

func DecryptWithPrivateKey(cipherText []byte) []byte {
	plaintext, err := PrivateKey.Decrypt(reader, cipherText, nil)

	if err != nil {
		panic(err)
	}

	return plaintext
}
