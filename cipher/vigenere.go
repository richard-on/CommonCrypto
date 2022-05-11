package cipher

import (
	"errors"
	"math"
)

type Vigenere struct {
	*Alphabet
	Key string
}

func (vi *Vigenere) Encrypt(input string) (encrypted string, encryptError error) {
	key, encryptError := vi.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	inputRune := []rune(input)
	var outputText []rune
	for i := 0; i < len(inputRune); i += len(key) {
		uIdx := math.Min(float64(i+len(key)), float64(len(inputRune)))
		oSlice := inputRune[i:int(uIdx)]

		var crS []rune
		for j := range oSlice {
			crS = append(crS, vi.symbols[(vi.SymbolIndex[oSlice[j]]+vi.SymbolIndex[key[j]])%vi.len])
		}
		outputText = append(outputText, crS...)
	}

	return string(outputText), nil
}

func (vi *Vigenere) Decrypt(input string) (decrypted string, decryptError error) {
	key, encryptError := vi.validateKey()
	if encryptError != nil {
		return "", decryptError
	}

	inputRune := []rune(input)
	var outputText []rune
	for i := 0; i < len(inputRune); i += len(key) {
		uIdx := math.Min(float64(i+len(key)), float64(len(inputRune)))
		oSlice := inputRune[i:int(uIdx)]

		var crS []rune
		for j := range oSlice {
			idx := (vi.SymbolIndex[oSlice[j]] - vi.SymbolIndex[key[j]]) % vi.len
			if idx < 0 {
				idx += vi.len
			}
			crS = append(crS, vi.symbols[idx])
		}
		outputText = append(outputText, crS...)
	}

	return string(outputText), nil
}

func (vi *Vigenere) validateKey() (keys []rune, validateError error) {
	runeKey := []rune(vi.Key)

	if len(runeKey) == 0 {
		return nil, errors.New("incorrect key: zero-length key")
	}

	for _, v := range vi.Key {
		if _, ok := vi.SymbolIndex[v]; !ok {
			return nil, errors.New("incorrect key: symbols not from alphabet")
		}
	}

	return runeKey, nil
}
