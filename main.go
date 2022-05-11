package main

import (
	"CryptoLab1/cipher"
	"fmt"
)

func EncryptAndDecrypt(c cipher.Cipher, input string) {
	fmt.Println(input)

	encrypted, err := c.Encrypt(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(encrypted)

	decrypted, err := c.Decrypt(encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println(decrypted)
}

func main() {
	engAlphabet, _ := cipher.NewAlphabet("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	rusAlphabet := cipher.DefaultAlphabet()

	trEng := cipher.Affine{
		Alphabet: &engAlphabet,
		Key:      "fr",
	}
	EncryptAndDecrypt(&trEng, "richard here")

	viRus := cipher.Vigenere{
		Alphabet: &rusAlphabet,
		Key:      "ЁЖИК",
	}
	EncryptAndDecrypt(&viRus, "ВЫШЕЛ ЁЖИК")

}
