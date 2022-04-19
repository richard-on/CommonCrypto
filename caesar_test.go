package main

import (
	"testing"
	"unicode/utf8"
)

func TestCaesar(t *testing.T) {
	testcases := []struct {
		alphabet, in, want string
		key                rune
	}{
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", "RICHARD", "RICHARD", 'D'},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ ", " ", " ", 'D'},
		{"АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ ", "РИЧАРД", "РИЧАРД", 'Д'},
	}
	for _, tc := range testcases {
		encrypt, _ := Encrypt(tc.alphabet, tc.in, tc.key)
		decrypt, _ := Decrypt(tc.alphabet, encrypt, tc.key)
		if decrypt != tc.want {
			t.Errorf("In: %v, encrypt: %v, decrypt: %v, want: %v", tc.in, encrypt, decrypt, tc.want)
		}
	}
}

func FuzzCaesar(f *testing.F) {
	testcases := []string{"RICHARD", " ", "ABCDEFG"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		encrypt, _ := Encrypt("ABCDEFGHIJKLMNOPQRSTUVWXYZ ", orig, 'D')
		decrypt, _ := Decrypt("ABCDEFGHIJKLMNOPQRSTUVWXYZ ", encrypt, 'D')
		if decrypt != orig {
			t.Errorf("In: %v, encrypt: %v, decrypt: %v", orig, encrypt, decrypt)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(decrypt) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", decrypt)
		}
	})

}
