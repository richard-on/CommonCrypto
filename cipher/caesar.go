package cipher

import (
	"errors"
	"strings"
)

type Caesar struct {
	*Alphabet
	Key string
}

func (c *Caesar) Encrypt(input string) (encrypted string, encryptError error) {
	key, encryptError := c.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	return c.rotate(input, key, Encrypt)
}

func (c *Caesar) Decrypt(input string) (decrypted string, decryptError error) {
	key, encryptError := c.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	return c.rotate(input, key, Decrypt)
}

func (c *Caesar) validateKey() (rune, error) {
	keyRune := []rune(c.Key)
	if len(keyRune) > 1 {
		return keyRune[0], errors.New("incorrect Key: not a single character")
	}
	if strings.Index(string(c.symbols), c.Key) == -1 {
		return keyRune[0], errors.New("incorrect Key: not in alphabet")
	}

	return keyRune[0], nil
}

func (c *Caesar) rotate(input string, key rune, operation Operation) (string, error) {
	inputRune := []rune(input)

	shift, _ := c.SymbolIndex[key]

	if operation == Decrypt {
		shift = -shift
	}

	text := make([]rune, len(inputRune))
	for i, v := range inputRune {
		idx, ok := c.SymbolIndex[v]
		if !ok {
			return "", errors.New("incorrect input: one or more characters not in alphabet")
		}

		pos := (idx + shift) % len(c.symbols)
		if pos < 0 {
			pos += c.len
		}

		text[i] = c.symbols[pos]
	}

	return string(text), nil
}
