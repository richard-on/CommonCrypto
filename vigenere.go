package main

import "errors"

type Vigenere struct {
	*Cipher
	key string
}

func (v *Vigenere) validateKey() (keys []rune, validateError error) {
	if len([]rune(v.key)) != len([]rune(v.alphabet)) {
		return nil, errors.New("incorrect key: not an alphabet length")
	}

	return []rune(v.key), nil
}

func (v *Vigenere) Encrypt(alphabet []rune, input string) (encrypted string, encryptError error) {
	keys, encryptError := v.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	return string(encryptedRune), nil
}

func (v *Vigenere) Decrypt(alphabet []rune, input string) (decrypted string, decryptError error) {
	keys, decryptError := v.validateKey()
	if decryptError != nil {
		return "", decryptError
	}

	return string(decryptedRune), nil
}
