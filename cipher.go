package main

import (
	"errors"
	"regexp"
)

type Operation int

const (
	encrypt Operation = iota
	decrypt
)

type CryptoSystem int

const (
	caesar CryptoSystem = iota
	affine
	simpleSubstitution
	hill
	transposition
	vigenere
)

type Alphabet struct {
	symbols     []rune
	symbolIndex map[rune]int
	intIndex    map[int]rune
}

type Cipher struct {
	CryptoSystem CryptoSystem
	alphabet     Alphabet
}

func (cipher *Cipher) Encrypt(input string, key string) (encrypted string, encryptError error) {
	alphabet, encryptError := cipher.validateAlphabet()
	if encryptError != nil {
		return "", encryptError
	}
	switch cipher.CryptoSystem {
	case caesar:
		caesar := Caesar{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := caesar.Encrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil

	case affine:
		affine := Affine{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := affine.Encrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil

	case simpleSubstitution:
		sub := Substitution{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := sub.Encrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil

	case hill:
		sub := Hill{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := sub.Encrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil

	case transposition:
		sub := Transposition{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := sub.Encrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil

	case vigenere:
		sub := Vigenere{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := sub.Encrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil
	}

	return encrypted, nil
}

func (cipher *Cipher) Decrypt(input string, key string) (decrypted string, decryptError error) {
	alphabet, decryptError := cipher.validateAlphabet()
	if decryptError != nil {
		return "", decryptError
	}
	switch cipher.CryptoSystem {
	case caesar:
		caesar := Caesar{
			Cipher: cipher,
			key:    key,
		}
		decrypted, err := caesar.Decrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return decrypted, nil

	case affine:
		affine := Affine{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := affine.Decrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil

	case simpleSubstitution:
		sub := Substitution{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := sub.Decrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil

	case hill:
		sub := Hill{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := sub.Decrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil

	case transposition:
		sub := Transposition{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := sub.Decrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil

	case vigenere:
		sub := Vigenere{
			Cipher: cipher,
			key:    key,
		}
		encrypted, err := sub.Decrypt(alphabet, input)
		if err != nil {
			return "", err
		}
		return encrypted, nil
	}

	return decrypted, nil
}

func (cipher *Cipher) validateAlphabet() ([]rune, error) {
	if match, err := regexp.MatchString("[^ а-яА-Яa-zA-Z\\p{P}\\p{S}]", string(cipher.alphabet.symbols)); err != nil || match {
		return nil, errors.New("incorrect alphabet: disallowed characters")
	}
	set := map[rune]bool{}
	alphabetRune := []rune(cipher.alphabet.symbols)
	for _, v := range alphabetRune {
		if _, ok := set[v]; ok {
			return nil, errors.New("incorrect alphabet: repeating symbols")
		}
		set[v] = true
	}

	return alphabetRune, nil
}
