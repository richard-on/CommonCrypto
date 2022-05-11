package cipher

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"strings"
)

type Hill struct {
	*Alphabet
	Key string
}

func (h *Hill) Encrypt(input string) (encrypted string, encryptError error) {
	keys, encryptError := h.validateKey()
	if encryptError != nil {
		return "", encryptError
	}

	inputRune := []rune(input)
	if len(inputRune)%2 != 0 {
		inputRune = append(inputRune, h.symbols[0])
	}

	var outputText strings.Builder

	keyFloat := make([]float64, 4)
	for i := 0; i < 4; i++ {
		keyFloat[i] = float64(keys[i].pos)
	}
	key := mat.NewDense(2, 2, keyFloat)

	for i := 0; i < len(inputRune)-1; i += 2 {

		vector := make([]float64, 2)
		for j, r := range inputRune[i : i+2] {
			vector[j] = float64(h.SymbolIndex[r])
		}

		vec := mat.NewVecDense(2, vector)
		vec.MulVec(key, vec)
		for j := 0; j < 2; j++ {
			r := h.intIndex[int(vec.AtVec(j))%h.len]
			outputText.WriteRune(r)

		}
	}

	return outputText.String(), nil
}

func (h *Hill) Decrypt(input string) (decrypted string, decryptError error) {
	keys, decryptError := h.validateKey()
	if decryptError != nil {
		return "", decryptError
	}

	inputRune := []rune(input)
	if len(inputRune)%2 != 0 {
		inputRune = append(inputRune, h.symbols[0])
	}

	var outputText strings.Builder

	keyFloat := make([]float64, 4)
	for i := 0; i < 4; i++ {
		keyFloat[i] = float64(keys[i].pos)
	}
	key := mat.NewDense(2, 2, keyFloat)

	err := key.Inverse(key)
	if err != nil {
		return "", errors.New("unable to inverse matrix")
	}

	for i := 0; i < len(inputRune)-1; i += 2 {

		vector := make([]float64, 2)
		for j, r := range inputRune[i : i+2] {
			vector[j] = float64(h.SymbolIndex[r])
		}

		vec := mat.NewVecDense(2, vector)
		vec.MulVec(key, vec)
		for j := 0; j < 2; j++ {
			r, ok := h.intIndex[int(vec.AtVec(j))%h.len]
			if ok {
				outputText.WriteRune(r)
			} else {
				outputText.WriteRune('?')
			}

		}
	}

	return outputText.String(), nil
}

/*func (h *Hill) operation(Key *mat.Dense, input []rune) (string, error) {
	var outputText strings.Builder


	err := Key.Inverse(Key)
	if err != nil {
		return "", errors.New("unable to inverse matrix")
	}

	for i := 0; i < len(input)-1; i += 2 {

		vector := make([]float64, 2)
		for j, r := range input[i : i+2] {
			vector[j] = float64(h.symbolIndex[r])
		}

		vec := mat.NewVecDense(2, vector)
		vec.MulVec(Key, vec)
		for j := 0; j < 2; j++ {
			r := h.intIndex[int(vec.AtVec(j))%h.len]
			outputText.WriteRune(r)

		}
	}

	return outputText.String(), nil
}*/

func (h *Hill) validateKey() ([]Keys, error) {
	var ok bool

	keyRune := []rune(h.Key)
	if len(keyRune) != 4 {
		return nil, errors.New("incorrect keys: not a 4 character sequence")
	}

	keys := make([]Keys, 4)
	for i, v := range keyRune {
		keys[i].key = v
		keys[i].pos, ok = h.SymbolIndex[v]
		if !ok {
			return nil, errors.New("incorrect keys: one or more not in alphabet")
		}
	}

	if (keys[0].pos*keys[3].pos-keys[1].pos*keys[2].pos)%h.len == 0 {
		return nil, errors.New("incorrect keys: (k11 * k22 - k12 * k21) % M == 0")
	} /*else if !coPrime(keys[0].pos*keys[3].pos-keys[1].pos*keys[2].pos, h.len) {
		return nil, errors.New("incorrect keys: (k11 * k22 - k12 * k21) and M should be co-prime")
	}*/

	return keys, nil
}

func (a *Alphabet) Contains(r rune) bool {
	_, found := a.SymbolIndex[r]
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
	return a.SymbolIndex[s], nil
}

func (a *Alphabet) Itos(i int) (rune, error) {
	r, found := a.intIndex[i]
	if !found {
		return 'x', fmt.Errorf("%d cannot be mapped to symbol", i)
	}
	return r, nil
}
