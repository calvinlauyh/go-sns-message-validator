package main

import (
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

func verify(data []byte, certData []byte, signature string) error {
	block, _ := pem.Decode(certData)
	if block == nil {
		return errors.New("X.509 certificate error")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return fmt.Errorf("Parse certificate error: %v\n", err)
	}

	decodeSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("Signature base64 decode error: %v\n", err)
	}

	err = cert.CheckSignature(x509.SHA1WithRSA, data, decodeSignature)
	return err
}

func main() {
	var certPath string
	var filePath string
	var signature string

	flag.StringVar(&certPath, "certPath", "_assets/fakecert.pem", "Path to the X.509 certificate")
	flag.StringVar(&filePath, "filePath", "", "Path to data file containing the signable data")
	flag.StringVar(&signature, "signature", "", "The signed string to verify")

	flag.Parse()

	data, err := readData(filePath)
	if err != nil {
		fmt.Printf("Read data error: %v\n", err)
		os.Exit(1)
	}

	certData, err := readData(certPath)
	if err != nil {
		fmt.Printf("Read certificate error: %v\n", err)
		os.Exit(1)
	}

	err = verify(data, certData, signature)
	if err == nil {
		fmt.Println("OK")
	} else {
		fmt.Println("Incorrect signature")
	}
}
