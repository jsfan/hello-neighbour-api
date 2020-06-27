package config

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func writePemEncoded(fileName string, contents []byte) error {
	fh, err := os.Open(fileName)
	if err != nil {
		return err
	}
	block := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   contents,
	}
	return pem.Encode(fh, &block)
}

func generateNewKeyPair(keyFiles *KeyPair) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	err = writePemEncoded(keyFiles.PrivateKey, x509.MarshalPKCS1PrivateKey(privateKey))
	if err != nil {
		return err
	}
	pubKey, err := x509.MarshalPKIXPublicKey(privateKey.PublicKey)
	if err != nil {
		return err
	}
	err = writePemEncoded(keyFiles.PublicKey, pubKey)
	if err != nil {
		return err
	}
	return nil
}

func readPemEncoded(fileName string) (keyBlock *pem.Block, errVal error) {
	pemKey, err := readFile(fileName)
	if err != nil {
		return nil, err
	}
	decodedKey, _ := pem.Decode(pemKey)
	if decodedKey == nil {
		return nil, errors.New(fmt.Sprintf(`Could not decode key "%s".`, fileName))
	}
	return decodedKey, nil
}

func readPrivateKey(keyFile string) (privKey *rsa.PrivateKey, errVal error) {
	privKeyRaw, err := readPemEncoded(keyFile)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(privKeyRaw.Bytes)
}

func readPublicKey(keyFile string) (pubKey *rsa.PublicKey, errVal error) {
	pubKeyRaw, err := readPemEncoded(keyFile)
	if err != nil {
		return nil, err
	}
	genericPubKey, err := x509.ParsePKIXPublicKey(pubKeyRaw.Bytes)
	if err != nil {
		return nil, err
	}
	pubKey, ok := genericPubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New(fmt.Sprintf(`Public key "%s" does not seem to be an RSA public key`, keyFile))
	}
	return pubKey, nil
}

func ReadKeyPair(keyFiles *KeyPair) (keyPair *rsa.PrivateKey, errVal error) {
	keyPair, err := readPrivateKey(keyFiles.PrivateKey)
	if err != nil {
		return nil, err
	}
	pubKey, err := readPublicKey(keyFiles.PublicKey)
	if err != nil {
		return nil, err
	}
	keyPair.PublicKey = *pubKey
	return keyPair, nil
}
