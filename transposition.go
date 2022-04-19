package main

import "errors"

type Transposition struct {
	*Cipher
	key string
}

func (t *Transposition) validateKey() (keys []rune, validateError error) {
	if len([]rune(t.key)) != len([]rune(t.alphabet)) {
		return nil, errors.New("incorrect key: not an alphabet length")
	}

	return []rune(t.key), nil
}

func (t *Transposition) Encrypt(alphabet []rune, input string) (encrypted string, encryptError error) {
	keys, encryptError := t.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	return string(encryptedRune), nil
}

func (t *Transposition) Decrypt(alphabet []rune, input string) (decrypted string, decryptError error) {
	keys, decryptError := t.validateKey()
	if decryptError != nil {
		return "", decryptError
	}

	return string(decryptedRune), nil
}
