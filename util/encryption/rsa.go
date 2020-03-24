package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
)

type EncryptionData struct {
	VerifyToken []byte
	DERPubkey   []byte
	RSAKey      *rsa.PrivateKey
}

var pubKey, privateKey = GetKeys()

var EncryptionDataInstance = EncryptionData{
	VerifyToken: GetVerifyToken(),
	DERPubkey:   pubKey,
	RSAKey:      privateKey,
}

var reader = rand.Reader

func GetVerifyToken() []byte {
	verifyTokenBytes := make([]byte, 4)
	_, _ = rand.Read(verifyTokenBytes)
	return verifyTokenBytes
}

func GetKeys() ([]byte, *rsa.PrivateKey) {
	key, err := rsa.GenerateKey(reader, 1024)
	checkError(err)

	asn1Bytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	checkError(err)

	return asn1Bytes, key
}

func checkError(err error) {
	if err != nil {
		log.Panic("Fatal error ", err.Error())
	}
}

func DecryptWithPrivateKey(cipherText []byte, key *rsa.PrivateKey) []byte {
	plaintext, err := key.Decrypt(reader, cipherText, nil)

	if err != nil {
		panic(err)
	}

	return plaintext
}
