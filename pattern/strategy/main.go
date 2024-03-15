package main

import "fmt"

type Encryptor interface {
	Encrypt(string) string
}

type EncryptionContext struct {
	Encryptor
}

func NewEnctyptionContext(e Encryptor) *EncryptionContext {
	return &EncryptionContext{Encryptor: e}
}

func (e *EncryptionContext) SwitchEncryptionAlgorythm(en Encryptor) {
	e.Encryptor = en
}

func main() {
	sha2 := &SHA256{}
	sha5 := &SHA512{}
	rsa := &RSA{}

	en := NewEnctyptionContext(sha2)
	fmt.Println(en.Encrypt("some data"))

	en.SwitchEncryptionAlgorythm(sha5)
	fmt.Println(en.Encrypt("some data"))

	en.SwitchEncryptionAlgorythm(rsa)
	fmt.Println(en.Encrypt("some data"))
}

type SHA256 struct{}

func (s SHA256) Encrypt(data string) string {
	return "encrypted by sha256"
}

type SHA512 struct{}

func (s SHA512) Encrypt(data string) string {
	return "encrypted by sha512"
}

type RSA struct{}

func (r RSA) Encrypt(data string) string {
	return "enctypted by rsa"
}
