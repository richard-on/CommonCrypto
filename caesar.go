package main

import (
	"errors"
	"strings"
)

type Caesar struct {
	*Cipher
	key string
}

func (caesar *Caesar) validateKey() (rune, error) {
	keyRune := []rune(caesar.key)
	if len(keyRune) > 1 {
		return keyRune[0], errors.New("incorrect key: not a single character")
	}
	if strings.Index(string(caesar.alphabet.symbols), caesar.key) == -1 {
		return keyRune[0], errors.New("incorrect key: not in alphabet")
	}

	return keyRune[0], nil
}

func (caesar *Caesar) rotate(alphabetRune []rune, input string, key rune, operation Operation) (string, error) {
	inputRune := []rune(input)
	var shift int
	for i, v := range alphabetRune {
		if v == key {
			shift = i
		}
	}
	if operation == decrypt {
		shift = -shift
	}

	runeText := make([]rune, len(inputRune))
	for i, val := range inputRune {
		alphabetIdx := -1
		for idx, v := range alphabetRune {
			if v == val {
				alphabetIdx = idx
			}
		}
		if alphabetIdx == -1 {
			return "", errors.New("input symbol is not in alphabet")
		}
		pos := (alphabetIdx + shift) % len(alphabetRune)
		if pos < 0 {
			pos += len(alphabetRune)
		}
		runeText[i] = alphabetRune[pos]
	}

	return string(runeText), nil
}

func (caesar *Caesar) Encrypt(alphabet []rune, input string) (encrypted string, encryptError error) {
	key, encryptError := caesar.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	if encrypted, encryptError = caesar.rotate(alphabet, input, key, encrypt); encryptError != nil {
		return "", encryptError
	}

	return encrypted, nil
}

func (caesar *Caesar) Decrypt(alphabet []rune, input string) (decrypted string, decryptError error) {
	key, encryptError := caesar.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	if decrypted, decryptError = caesar.rotate(alphabet, input, key, decrypt); decryptError != nil {
		return "", encryptError
	}

	return decrypted, nil
}
