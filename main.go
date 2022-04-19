package main

import (
	"fmt"
)

func main() {
	inputEng := "ПРАВАЯ"
	/*affineEng := Cipher{
		CryptoSystem: affine,
		alphabet:     " ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}*/

	subEng := Cipher{
		CryptoSystem: simpleSubstitution,
		alphabet:     " АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ",
	}

	encrypted, err := subEng.Encrypt(inputEng, " ЯБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮА")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	decrypted, err := subEng.Decrypt(encrypted, " ЯБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮА")
	if err != nil {
		return
	}

	fmt.Println(inputEng)
	fmt.Println(encrypted)
	fmt.Println(decrypted)
}
