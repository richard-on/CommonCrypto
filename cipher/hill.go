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

/*func residue(a, m int) float64 {
	reminder := int(math.Abs(float64(a % m)))
	if a >= 0 {
		return float64(reminder)
	} else if a < 0 && reminder != 0 {
		return float64(m - reminder)
	}
	return 0
}

func modularInverse(a, m int) (int, error) {
	if m == 1 {
		return 0, nil
	}
	if _, _, g := egcd(a, m); g != 1 {
		return -1, fmt.Errorf("%d and %d are not coprimes", a, m)
	}
	m0, x, y := m, 1, 0
	for a > 1 {
		q := a / m
		m, a = a%m, m
		y, x = x-q*y, y
	}
	if x < 0 {
		x += m0
	}
	return x, nil
}

func cofactor(key *mat.Dense) (*mat.Dense, error) {
	//cof := &Matrix{order: m.order, data: make([][]int, m.order)}
	for i := 0; i < 2; i++ {
		row := make([]int, 2)
		for j := 0; j < 2; j++ {
			minor, _ := Minor(m, i, j) // Error is neglected since row & col are always in bound
			detM, err := mat.Det(minor)
			if err != nil {
				return nil, fmt.Errorf("failed to compute det(m) for minor at row:%d col:%d\n%s;%v", i, j, m, err)
			}
			if (i+j)%2 == 0 {
				row[j] = detM
			} else {
				row[j] = detM * -1
			}
		}
		cof.data[i] = row
	}
	return cof, nil
}

func minor(key *mat.Dense, p, q int) (*mat.Dense, error) {
	r := mat.NewDense(1, 1, 0)
	r := &Matrix{order: m.order - 1}
	r.data = make([][]int, 0, r.order)
	for row := 0; row < m.order; row++ {
		tmp := make([]int, 0, r.order)
		for col := 0; col < m.order; col++ {
			if row != p && col != q {
				tmp = append(tmp, m.data[row][col])
			}
		}
		if row != p {
			r.data = append(r.data, tmp)
		}
	}
	return r, nil
}

func adJoint(key *mat.Dense) (*mat.Dense, error) {
	cof, err := key.Cofactor()
	if err != nil {
		return nil, fmt.Errorf("failed to compute cofactor matrix for \n%s; %v", m, err)
	}
	return cof.Transpose(), nil
}

func inverseMod(key *mat.Dense) (*mat.Dense, error) {
	det := mat.Det(key)
	res := residue(int(det), 2)
	inverse, _ := modularInverse(int(res), 2)
	adj, err := adJoint(key)
	if err != nil {
		return nil, fmt.Errorf("failed to compute Adj(\n%v\n); %v", key, err)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			adj.Set(i, j, residue(int(adj.At(i, j)*float64(inverse)), 2))
		}
	}

	return adj, nil
}*/

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
