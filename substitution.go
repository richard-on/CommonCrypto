package main

import (
	"errors"
)

type Substitution struct {
	*Cipher
	key string
}

func (s *Substitution) validateKey() (keys []rune, validateError error) {
	if len([]rune(s.key)) != len([]rune(s.alphabet.symbols)) {
		return nil, errors.New("incorrect key: not an alphabet length")
	}

	set := map[rune]bool{}
	alphabetRune := []rune(s.alphabet.symbols)
	for _, v := range alphabetRune {
		if _, ok := set[v]; ok {
			return nil, errors.New("incorrect key: repeating symbols")
		}
		set[v] = true
	}

	for _, v := range set {
		if !v {
			return nil, errors.New("incorrect key: unrepresented symbols")
		}
	}

	return []rune(s.key), nil
}

func (s *Substitution) Encrypt(alphabet []rune, input string) (encrypted string, encryptError error) {
	keys, encryptError := s.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	encryptedRune := make([]rune, len([]rune(input)))
	symbolIdx := 0
	for i, inputSymbol := range []rune(input) {
		for j, symbol := range alphabet {
			if symbol == inputSymbol {
				symbolIdx = j
			}
		}
		encryptedRune[i] = keys[symbolIdx]

	}

	return string(encryptedRune), nil
}

func (s *Substitution) Decrypt(alphabet []rune, input string) (decrypted string, decryptError error) {
	keys, decryptError := s.validateKey()
	if decryptError != nil {
		return "", decryptError
	}

	decryptedRune := make([]rune, len([]rune(input)))
	symbolIdx := 0
	for i, inputSymbol := range []rune(input) {
		for j, symbol := range keys {
			if symbol == inputSymbol {
				symbolIdx = j
			}
		}
		decryptedRune[i] = alphabet[symbolIdx]

	}

	return string(decryptedRune), nil
}
