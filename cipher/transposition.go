package cipher

import (
	"errors"
	"sort"
)

type Transposition struct {
	*Alphabet
	Key string
}

func (t *Transposition) Encrypt(input string) (encrypted string, encryptError error) {
	key, encryptError := t.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	return t.rotate(input, key, Encrypt)
}

func (t *Transposition) Decrypt(input string) (decrypted string, decryptError error) {
	key, decryptError := t.validateKey()
	if decryptError != nil {
		return "", decryptError
	}

	return t.rotate(input, key, Decrypt)
}

func (t *Transposition) rotate(input string, key []rune, operation Operation) (string, error) {
	inputRune := []rune(input)
	permutations := getKey(key)

	keyLength := len(key)
	textLength := len(inputRune)
	for i := 0; i < len(key)-textLength%keyLength; i++ {
		inputRune = append(inputRune, t.symbols[t.len-1])
	}

	result := make([]rune, len(inputRune))
	for i := 0; i < len(inputRune); i += len(key) {
		transposition := make([]rune, len(key))
		for j := 0; j < len(key); j++ {
			if operation == Encrypt {
				transposition[permutations[j]-1] = inputRune[i+j]
			} else {
				transposition[j] = inputRune[i+permutations[j]-1]
			}
		}
		result = append(result, transposition...)
	}

	return string(result), nil
}

func (t *Transposition) validateKey() (keys []rune, validateError error) {
	runeKey := []rune(t.Key)

	if len(runeKey) > t.len {
		return nil, errors.New("incorrect key: more than alphabet length")
	}

	set := map[rune]bool{}
	for _, v := range t.Key {
		if _, ok := set[v]; ok {
			return nil, errors.New("incorrect key: repeating symbols")
		}
		if _, ok := t.SymbolIndex[v]; !ok {
			return nil, errors.New("incorrect key: symbols not from alphabet")
		}
		set[v] = true
	}

	return runeKey, nil
}

func getKey(key []rune) []int {
	var sortedKey = make([]rune, len(key))
	copy(sortedKey, key)

	sort.Slice(sortedKey, func(i, j int) bool { return sortedKey[i] < sortedKey[j] })

	usedLettersMap := make(map[rune]int)
	wordLength := len(key)
	resultKey := make([]int, wordLength)

	for i := 0; i < len(key); i++ {
		symbol := key[i]
		numberOfUsage := usedLettersMap[symbol]
		resultKey[i] = getIndex(sortedKey, symbol) + numberOfUsage + 1
		numberOfUsage++
		usedLettersMap[symbol] = numberOfUsage
	}

	return resultKey
}

func getIndex(wordSet []rune, subString rune) int {
	for i, v := range wordSet {
		if v == subString {
			return i
		}
	}

	return -1
}
