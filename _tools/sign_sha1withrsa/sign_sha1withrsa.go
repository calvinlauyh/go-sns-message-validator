package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Read the data from file into slice of bytes
func readData(path string) ([]byte, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	return ioutil.ReadAll(fp)
}

func sign(data []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("RSA private key error")
	}
	if block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("Provided key is not RSA private key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Parse private key error: %v", err)
	}

	h := sha1.New()
	h.Write(data)
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, hashed)
	if err != nil {
		return nil, fmt.Errorf("Sign error: %v", err)
	}

	return signature, nil
}

func main() {
	var keyPath string
	var filePath string
	var base64Encode bool

	flag.StringVar(&keyPath, "keyPath", "_assets/fakecert.key", "Path to the prviate key")
	flag.StringVar(&filePath, "filePath", "", "Path to data file containing the signable data")
	flag.BoolVar(&base64Encode, "base64Encode", true, "Whether to output base64 encoded signature")

	flag.Parse()

	data, err := readData(filePath)
	if err != nil {
		fmt.Printf("Read data error: %v\n", err)
		os.Exit(1)
	}

	key, err := readData(keyPath)
	if err != nil {
		fmt.Printf("Read key error: %v\n", err)
		os.Exit(1)
	}

	signature, err := sign(data, key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if base64Encode {
		fmt.Println(base64.StdEncoding.EncodeToString(signature))
	} else {
		fmt.Println(signature)
	}
}
