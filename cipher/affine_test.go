package cipher

import (
	"testing"
)

func TestAffine(t *testing.T) {
	testcases := []struct {
		alphabetStr, key, in, want string
	}{
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "DH", "RICHARD HERE", "RICHARD HERE"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "BF", "ABCDE", "ABCDE"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "HC", "SUNNY  DAY", "SUNNY  DAY"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "WW", "JK KJEVV KHEFVBKEBVEVBKEBV", "JK KJEVV KHEFVBKEBVEVBKEBV"},
		{"abcdefghijklmnopqrstuvwxyz", "el", "richardhere", "richardhere"},
		{"АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ ", "ДБ", "РИЧАРД ТУТ", "РИЧАРД ТУТ"},
	}

	for _, tc := range testcases {
		alphabet, _ := NewAlphabet(tc.alphabetStr)
		caesar := Affine{
			Alphabet: &alphabet,
			Key:      tc.key,
		}
		enc, _ := caesar.Encrypt(tc.in)
		dec, _ := caesar.Decrypt(enc)
		if dec != tc.want {
			t.Errorf("In: %v, encrypt: %v, decrypt: %v, want: %v", tc.in, enc, dec, tc.want)
		}
	}
}

func FuzzAffine(f *testing.F) {
	testcases := []struct {
		alphabetStr, key, orig string
	}{
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "DH", "RICHARD HERE"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "BF", "ABCDE"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "HC", "SUNNY  DAY"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "WW", "JK KJEVV KHEFVBKEBVEVBKEBV"},
		{"abcdefghijklmnopqrstuvwxyz", "el", "richardhere"},
		{"АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ ", "ДБ", "РИЧАРД ТУТ"},
	}

	for _, tc := range testcases {
		f.Add(tc.key, tc.orig)
	}

	f.Fuzz(func(t *testing.T, key, orig string) {
		alphabet, _ := NewAlphabet("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
		affine := Affine{
			Alphabet: &alphabet,
			Key:      key,
		}

		if len(key) != 2 {
			t.Skip()
		}
		if _, ok := alphabet.SymbolIndex[[]rune(key)[0]]; !ok {
			t.Skip()
		}

		keys := make([]int, 2)
		var ok bool
		for i, v := range []rune(key) {
			keys[i], ok = alphabet.SymbolIndex[v]
			if !ok {
				t.Skip()
			}
			keys[i]++
		}
		if !coPrime(keys[0], affine.len) {
			t.Skip()
		}
		for _, v := range []rune(orig) {
			if _, ok := alphabet.SymbolIndex[v]; !ok {
				t.Skip()
			}
		}

		enc, _ := affine.Encrypt(orig)
		dec, _ := affine.Decrypt(enc)
		if dec != orig {
			t.Errorf("Original input: %v, encrypted: %v, decrypted: %v, key: %v", orig, enc, dec, key)
		}
	})
}
