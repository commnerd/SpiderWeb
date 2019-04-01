package main

import (
    "golang.org/x/crypto/ssh"
    "encoding/pem"
    "crypto/x509"
    "crypto/rand"
    "crypto/rsa"
    "io/ioutil"
    "log"
    "os"
)

func GenerateKeys() {
    bitSize := 4096

    privateKey, err := generatePrivateKey(bitSize)
    if err != nil {
        panic(err)
    }

    publicKeyBytes, err := generatePublicKey(&privateKey.PublicKey)
    if err != nil {
        panic(err)
    }

    privateKeyBytes, err := encodePrivateKeyToPEM(privateKey)
    if err != nil {
        panic(err)
    }

    writeToFiles(publicKeyBytes, privateKeyBytes)
}

// generatePrivateKey creates an RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
    // Private Key generation
    privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
    if err != nil {
        return nil, err
    }

    // Validate Private Key
    err = privateKey.Validate()
    if err != nil {
        return nil, err
    }

    log.Println("Private Key generated")
    return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) ([]byte, error) {
    // Get ASN.1 DER format
    privDER := x509.MarshalPKCS1PrivateKey(privateKey)

    // pem.Block
    privBlock := pem.Block{
        Type:    "RSA PRIVATE KEY",
        Headers: nil,
        Bytes:   privDER,
    }

    // Private key in PEM format
    privatePEM := pem.EncodeToMemory(&privBlock)

    return privatePEM, nil
}

// generatePublicKey take an rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format 'ssh-rsa ...'
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
    publicRsaKey, err := ssh.NewPublicKey(privatekey)
    if err != nil {
        return nil, err
    }

    pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

    log.Println("Public key generated")
    return pubKeyBytes, nil
}

func writeToFiles(publicKey []byte, privateKey []byte) {
    var path string = "/root/.ssh"
    if _, err := os.Stat(path); os.IsNotExist(err) {
        os.Mkdir(path, 0755)
    }

    err := ioutil.WriteFile("/root/.ssh/id_rsa.pub", publicKey, 0400)
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile("/root/.ssh/id_rsa", privateKey, 0400)
    if err != nil {
        panic(err)
    }
}
