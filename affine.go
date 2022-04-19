package main

import (
	"errors"
	"strings"
)

type Affine struct {
	*Cipher
	key string
}

type Keys struct {
	key rune
	pos int
}

func coPrime(a, b int) bool {
	var t int
	for b != 0 {
		t = b
		b = a % b
		a = t
	}

	return a == 1
}

func egcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	g, y, x := egcd(b%a, a)
	return g, x - (b/a)*y, y
}

func inv(a, m int) int {
	_, x, _ := egcd(a, m)

	return x % m
}

func (affine *Affine) validateKey() ([]Keys, error) {
	keys := make([]Keys, 2)
	keyRune := []rune(affine.key)
	if len(keyRune) != 2 {
		return nil, errors.New("incorrect keys: not a two character sequence")
	}
	for i, v := range keyRune {
		keys[i].key = v
		keys[i].pos = strings.IndexRune(string(affine.alphabet.symbols), keyRune[i])
		if keys[i].pos == -1 {
			return nil, errors.New("incorrect keys: one or more not in alphabet")
		}
	}
	if !coPrime(keys[0].pos, len(affine.alphabet.symbols)) {
		return keys, errors.New("incorrect key 0: not co-prime with alphabet length")
	}

	return keys, nil
}

func (affine *Affine) rotate(alphabetRune []rune, input string, keys []Keys, operation Operation) (string, error) {
	inputRune := []rune(input)
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
		var pos int
		if operation == encrypt {
			pos = (keys[0].pos*alphabetIdx + keys[1].pos) % len(alphabetRune)
		} else {
			pos = (inv(keys[0].pos, len(alphabetRune))) * (alphabetIdx - keys[1].pos) % len(alphabetRune)
		}

		if pos < 0 {
			pos += len(alphabetRune)
		}
		runeText[i] = alphabetRune[pos]
	}

	return string(runeText), nil
}

func (affine *Affine) Encrypt(alphabet []rune, input string) (encrypted string, encryptError error) {
	key, encryptError := affine.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	if encrypted, encryptError = affine.rotate(alphabet, input, key, encrypt); encryptError != nil {
		return "", encryptError
	}

	return encrypted, nil
}

func (affine *Affine) Decrypt(alphabet []rune, input string) (decrypted string, decryptError error) {
	key, encryptError := affine.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	if decrypted, decryptError = affine.rotate(alphabet, input, key, decrypt); decryptError != nil {
		return "", encryptError
	}

	return decrypted, nil
}
