package cipher

import (
	"errors"
)

type Affine struct {
	*Alphabet
	Key string
}

func (a *Affine) Encrypt(input string) (encrypted string, encryptError error) {
	key, encryptError := a.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	if encrypted, encryptError = a.rotate(input, key, Encrypt); encryptError != nil {
		return "", encryptError
	}

	return encrypted, nil
}

func (a *Affine) Decrypt(input string) (decrypted string, decryptError error) {
	key, encryptError := a.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	if decrypted, decryptError = a.rotate(input, key, Decrypt); decryptError != nil {
		return "", encryptError
	}

	return decrypted, nil
}

func (a *Affine) validateKey() ([]Keys, error) {
	var ok bool
	keys := make([]Keys, 2)
	keyRune := []rune(a.Key)
	if len(keyRune) != 2 {
		return nil, errors.New("incorrect keys: not a two character sequence")
	}

	for i, v := range keyRune {
		keys[i].key = v
		keys[i].pos, ok = a.SymbolIndex[v]
		if !ok {
			return nil, errors.New("incorrect keys: one or more not in alphabet")
		}
		keys[i].pos++
	}

	if !coPrime(keys[0].pos, a.len) {
		return nil, errors.New("incorrect Key 0: not co-prime with alphabet length")
	}

	return keys, nil
}

func (a *Affine) rotate(input string, keys []Keys, operation Operation) (string, error) {
	inputRune := []rune(input)

	outputText := make([]rune, len(inputRune))
	for i, v := range inputRune {
		_, ok := a.SymbolIndex[v]
		if !ok {
			return "", errors.New("incorrect input: one or more characters not in alphabet")
		}

		var pos int
		if operation == Encrypt {
			pos = (keys[0].pos*a.SymbolIndex[v] + keys[1].pos) % a.len
		} else {
			pos = (inv(keys[0].pos, a.len)) * (a.SymbolIndex[v] - keys[1].pos) % a.len
		}

		if pos < 0 {
			pos += a.len
		}
		outputText[i] = a.symbols[pos]
	}

	return string(outputText), nil
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
