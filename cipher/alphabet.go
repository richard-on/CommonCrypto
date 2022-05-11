package cipher

import (
	"os"
	"regexp"
)

const (
	defaultAlphabet = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ "
)

type Alphabet struct {
	len         int
	str         string
	symbols     []rune
	SymbolIndex map[rune]int
	intIndex    map[int]rune
}

func NewAlphabet(alphabetString string) (Alphabet, error) {
	var a Alphabet
	var err error

	if alphabetString == "" {
		return DefaultAlphabet(), nil

	} else if a.symbols, err = a.validateAlphabet(alphabetString); err != nil {
		return Alphabet{}, err
	} else {
		a.str = alphabetString
	}

	a.len = len(a.symbols)
	a.SymbolIndex = make(map[rune]int, len(a.symbols))
	a.intIndex = make(map[int]rune, len(a.symbols))
	for i, v := range a.symbols {
		a.SymbolIndex[v] = i
		a.intIndex[i] = v
	}

	return a, nil
}

func NewAlphabetFromFile(fileName string) (Alphabet, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	if fileName == "" {
		return DefaultAlphabet(), nil
	}

	return NewAlphabet(string(data))
}

func DefaultAlphabet() Alphabet {
	var a Alphabet

	a.str = defaultAlphabet
	a.symbols = []rune(defaultAlphabet)
	a.len = len(a.symbols)

	a.SymbolIndex = make(map[rune]int, len(a.symbols))
	a.intIndex = make(map[int]rune, len(a.symbols))
	for i, v := range a.symbols {
		a.SymbolIndex[v] = i
		a.intIndex[i] = v
	}

	return a
}

func (a *Alphabet) validateAlphabet(alphabetString string) ([]rune, error) {
	if match, err := regexp.MatchString("[^ [а-яА-Я][a-zA-Z]\\p{P}\\p{S}]", alphabetString); err != nil || match {
		panic("incorrect alphabet: disallowed characters")
	}

	set := make(map[rune]bool, len(a.symbols))
	alphabetRune := []rune(alphabetString)
	for _, v := range alphabetRune {
		if _, ok := set[v]; ok {
			panic("incorrect alphabet: repeating symbols")
		}
		set[v] = true
	}

	return alphabetRune, nil
}
