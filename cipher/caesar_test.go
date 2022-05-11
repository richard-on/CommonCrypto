package cipher

import (
	"testing"
)

func TestCaesar(t *testing.T) {
	testcases := []struct {
		alphabetStr, key, in, want string
	}{
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "D", "RICHARD HERE", "RICHARD HERE"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "B", "ABCDE", "ABCDE"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", " ", "SUNNY  DAY", "SUNNY  DAY"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "W", "JK KJEVV KHEFVBKEBVEVBKEBV", "JK KJEVV KHEFVBKEBVEVBKEBV"},
		{"abcdefghijklmnopqrstuvwxyz", "e", "richardhere", "richardhere"},
		{"АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ ", "Д", "РИЧАРД ТУТ", "РИЧАРД ТУТ"},
	}

	for _, tc := range testcases {
		alphabet, _ := NewAlphabet(tc.alphabetStr)
		caesar := Caesar{
			Alphabet: &alphabet,
			Key:      tc.key,
		}
		enc, _ := caesar.Encrypt(tc.in)
		dec, _ := caesar.Decrypt(enc)
		if dec != tc.want {
			t.Errorf("In: %v, encrypt: %v, decrypt: %v, want: %v", tc.in, Encrypt, Decrypt, tc.want)
		}
	}
}

func FuzzCaesar(f *testing.F) {
	testcases := []struct {
		alphabetStr, key, orig string
	}{
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "D", "RICHARD HERE"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "B", "ABCDE"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", " ", "SUNNY  DAY"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "W", "JK KJEVV KHEFVBKEBVEVBKEBV"},
	}

	for _, tc := range testcases {
		f.Add(tc.key, tc.orig)
	}

	f.Fuzz(func(t *testing.T, key, orig string) {
		alphabet, _ := NewAlphabet("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
		caesar := Caesar{
			Alphabet: &alphabet,
			Key:      key,
		}

		if len(key) != 1 {
			t.Skip()
		}
		if _, ok := alphabet.SymbolIndex[[]rune(key)[0]]; !ok {
			t.Skip()
		}
		for _, v := range []rune(orig) {
			if _, ok := alphabet.SymbolIndex[v]; !ok {
				t.Skip()
			}
		}

		enc, _ := caesar.Encrypt(orig)
		dec, _ := caesar.Decrypt(enc)
		if dec != orig {
			t.Errorf("Original input: %v, encrypted: %v, decrypted: %v, key: %v", orig, enc, dec, key)
		}
	})
}
