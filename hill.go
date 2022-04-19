package main

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/big"
)

type Hill struct {
	*Cipher
	key string
}

func (a *Alphabet) Contains(r rune) bool {
	_, found := a.symbolIndex[r]
	return found
}

func (a *Alphabet) Belongs(s string) bool {
	for _, r := range []rune(s) {
		if !a.Contains(r) {
			return false
		}
	}
	return true
}

func (a *Alphabet) Stoi(s rune) (int, error) {
	if !a.Contains(s) {
		return -1, fmt.Errorf("symbols %q is not part of the alphabet", s)
	}
	return a.symbolIndex[s], nil
}

// Itos returns the symbol value of the given int i (Int To Symbol).
func (a *Alphabet) Itos(i int) (rune, error) {
	r, found := a.intIndex[i]
	if !found {
		return 'x', fmt.Errorf("%d cannot be mapped to symbol", i)
	}
	return r, nil
}

func (hill *Hill) validateKey() (keys []rune, validateError error) {
	keys = make([]rune, 4)
	keyRune := []rune(hill.key)
	if len(keyRune) != 4 {
		return nil, errors.New("incorrect keys: not a 4 character sequence")
	}
	if !hill.alphabet.Belongs(string(keys)) {
		return nil, fmt.Errorf("message %q does not belong to alphabet %q", keys, hill.alphabet)
	}
	/*if !hill.alphabet.Belongs(rawK) {
		return nil, fmt.Errorf("key %q does not belong to alphabet %q", rawK, c.alphabet)
	}*/
	k := keys
	kInt := make([]int, len(k))
	for i, s := range k {
		kInt[i], _ = hill.alphabet.Stoi(s)
	}

	return keys, nil
}

func (hill *Hill) Encrypt(alphabet []rune, input string) (encrypted string, encryptError error) {
	keys, encryptError := hill.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	inputRune := []rune(input)
	keyFloat := make([]float64, 4)
	for i := 0; i < 4; i++ {
		keyFloat[i] = float64(keys[i])
	}
	matrix := mat.NewDense(2, 2, keyFloat)
	det := mat.Det(matrix)
	gcd := new(big.Int)

	if gcd.GCD(big.NewInt(int64(det)), big.NewInt(int64(len(alphabet))), big.NewInt(0), big.NewInt(0)) != big.NewInt(1) {
		return "", errors.New("incorrect matrix")
	}

	if len(inputRune)%2 != 0 {
		inputRune = append(inputRune, alphabet[0])
	}

	encryptedRune := make([]rune, len(inputRune))
	for i := 0; i < len(inputRune)-1; i += 2 {

		vecFloat := make([]float64, 2)
		//symbolIdx := 0
		for x := 0; x < len(inputRune); x += 2 {
			for j, symbol := range alphabet {
				if symbol == inputRune[x] {
					vecFloat = append(vecFloat, float64(j))
				}
			}
		}
		vec := mat.NewDense(2, 1, vecFloat)
		crypted := mat.NewDense(2, 1, vecFloat)

		crypted.Mul(vec, matrix)
		//crypted %= len(alphabet)

	}

	return string(encryptedRune), nil
}

func (hill *Hill) Decrypt(alphabet []rune, input string) (decrypted string, decryptError error) {
	keys, decryptError := hill.validateKey()
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
