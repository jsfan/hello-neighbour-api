package config

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func writePemEncoded(fileName string, keyType string, contents []byte) error {
	fh, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	block := pem.Block{
		Type:    keyType,
		Headers: nil,
		Bytes:   contents,
	}
	return pem.Encode(fh, &block)
}

func generateNewKeyPair(keyFiles *KeyPair) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	err = writePemEncoded(keyFiles.PrivateKey, "RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(privateKey))
	if err != nil {
		return err
	}
	pubKey, err := asn1.Marshal(privateKey.PublicKey)
	if err != nil {
		return err
	}
	err = writePemEncoded(keyFiles.PublicKey, "RSA PUBLIC KEY", pubKey)
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
	pubKey, err = x509.ParsePKCS1PublicKey(pubKeyRaw.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}

func ReadKeyPair(keyFiles *KeyPair) (keyPair *rsa.PrivateKey, errVal error) {
	if _, err := os.Stat(keyFiles.PrivateKey); err != nil && os.IsNotExist(err) {
		if err := generateNewKeyPair(keyFiles); err != nil {
			return nil, errors.Wrap(err, "could not generate new signing key pair")
		}
	}
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
