package cipher

import (
	"errors"
)

type Substitution struct {
	*Alphabet
	Key string
}

func (s *Substitution) Encrypt(input string) (encrypted string, encryptError error) {
	key, encryptError := s.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	inputRune := []rune(input)

	outputText := make([]rune, len([]rune(input)))

	for i, v := range inputRune {
		curSymbol, ok := s.SymbolIndex[v]
		if ok {
			outputText[i] = key[curSymbol]
		}
	}

	return string(outputText), nil
}

func (s *Substitution) Decrypt(input string) (decrypted string, decryptError error) {
	key, decryptError := s.validateKey()
	if decryptError != nil {
		return "", decryptError
	}

	inputRune := []rune(input)

	outputText := make([]rune, len([]rune(input)))

	symbolIdx := 0
	for i, v := range inputRune {
		if _, ok := s.SymbolIndex[v]; !ok {
			return "", errors.New("incorrect input: one or more characters not in alphabet")
		}

		for j, curSymbol := range key {
			if curSymbol == v {
				symbolIdx = j
			}
		}
		outputText[i] = s.symbols[symbolIdx]
	}

	return string(outputText), nil
}

func (s *Substitution) validateKey() (key []rune, validateError error) {
	if len([]rune(s.Key)) != s.len {
		return nil, errors.New("incorrect Key: not an alphabet length")
	}

	set := map[rune]bool{}
	for _, v := range s.Key {
		if _, ok := set[v]; ok {
			return nil, errors.New("incorrect Key: repeating symbols")
		}
		if _, ok := s.SymbolIndex[v]; !ok {
			return nil, errors.New("incorrect Key: symbols not from alphabet")
		}
		set[v] = true
	}

	return []rune(s.Key), nil
}
