package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
)

// Hash compute hash of the input string
func Hash(data string) []byte {
	h := sha1.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)
	return bs
}

//Encode64 encode the input byte array to base64 string
func Encode64(in []byte) string {
	return base64.StdEncoding.EncodeToString(in)
}

//Decode64 decode the base64 input string as byte array
func Decode64(in string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(in)
}

// GenerateKeyPair generates a new key pair
func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("error %+v\n", err)
	}
	return privkey, &privkey.PublicKey
}

// PrivateKeyToString convert the input private key to base64 string
func PrivateKeyToString(priv *rsa.PrivateKey) string {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return Encode64(privBytes)
}

// PublicKeyToString convert the input public key to base64 string
func PublicKeyToString(pub *rsa.PublicKey) string {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		log.Fatalf("error %+v\n", err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return Encode64(pubBytes)
}

// StringToPrivateKey convert the input base64 string to private key
func StringToPrivateKey(privS string) *rsa.PrivateKey {
	priv, _ := Decode64(privS)
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			log.Fatalf("error %+v\n", err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		log.Fatalf("error %+v\n", err)
	}
	return key
}

// StringToPublicKey convert the input base64 string to public key
func StringToPublicKey(pubS string) *rsa.PublicKey {
	pub, _ := Decode64(pubS)
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			log.Fatalf("error %+v\n", err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		log.Fatalf("error %+v\n", err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		log.Fatalf("error NotOK\n")
	}
	return key
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(data []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, data, nil)
	if err != nil {
		log.Fatalf("error %+v\n", err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		log.Fatalf("error %+v\n", err)
	}
	return plaintext
}

// Sign sign the input data with the given private key
func Sign(data []byte, pk *rsa.PrivateKey) (string, error) {
	signed, err := rsa.SignPKCS1v15(rand.Reader, pk, crypto.SHA256, data)
	if err != nil {
		return "", err
	}
	return Encode64(signed), nil
}

// VerifySign check the input signed string validation with the given data and public key
func VerifySign(data []byte, sigS string, pk *rsa.PublicKey) error {
	sig, _ := Decode64(sigS)
	return rsa.VerifyPKCS1v15(pk, crypto.SHA256, data, sig)
}
