package examples

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

func Run() {
	engAlphabet, _ := cipher.NewAlphabet("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	//rusAlphabet := cipher.DefaultAlphabet()

	fmt.Println("Caesar")
	EncryptAndDecrypt(&cipher.Caesar{
		Alphabet: &engAlphabet,
		Key:      "f",
	},
		"richard here")
	fmt.Printf("-----------------\n")

	fmt.Println("Affine")
	EncryptAndDecrypt(&cipher.Affine{
		Alphabet: &engAlphabet,
		Key:      "fh",
	},
		"richard here")
	fmt.Printf("-----------------\n")

	fmt.Println("Substitution")
	EncryptAndDecrypt(&cipher.Substitution{
		Alphabet: &engAlphabet,
		Key:      "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 ",
	},
		"richard here")
	fmt.Printf("-----------------\n")

	fmt.Println("Hill")
	EncryptAndDecrypt(&cipher.Hill{
		Alphabet: &engAlphabet,
		Key:      "fgrs",
	},
		"richard here")
	fmt.Printf("-----------------\n")

	fmt.Println("Transposition")
	EncryptAndDecrypt(&cipher.Transposition{
		Alphabet: &engAlphabet,
		Key:      "fdgwervlcz",
	},
		"richard here")
	fmt.Printf("-----------------\n")

	fmt.Println("Vigenere")
	EncryptAndDecrypt(&cipher.Vigenere{
		Alphabet: &engAlphabet,
		Key:      "fdfgdfwrhernjulty",
	},
		"richard here")
	fmt.Printf("-----------------\n")
}
