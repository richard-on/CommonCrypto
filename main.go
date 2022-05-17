package main

import (
	"CryptoLab1/cipher"
	"flag"
	"fmt"
	"os"
)

//engAlphabet, _ := cipher.NewAlphabet("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
//rusAlphabet := cipher.DefaultAlphabet()

func main() {
	//examples.Run()

	var alphabetFileName, inputFileName, keyFileName, cryptoSystem, operation string
	var c cipher.Cipher
	var alphabet cipher.Alphabet

	flag.StringVar(&alphabetFileName, "a", "", "alphabet filename")
	flag.StringVar(&inputFileName, "i", "in.txt", "input filename")
	flag.StringVar(&keyFileName, "k", "key.txt", "key filename")
	flag.StringVar(&cryptoSystem, "c", "c", "crypto system")
	flag.StringVar(&operation, "o", "e", "operation")
	flag.Parse()

	if alphabetFileName == "" {
		alphabet = cipher.DefaultAlphabet()
	} else {
		var err error
		alphabet, err = cipher.NewAlphabetFromFile(alphabetFileName)
		if err != nil {
			panic(err)
		}
	}

	in, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	key, err := os.ReadFile(keyFileName)
	if err != nil {
		panic(err)
	}

	switch cryptoSystem {
	case "a":
		c = &cipher.Affine{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	case "c":
		c = &cipher.Caesar{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	case "h":
		c = &cipher.Hill{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	case "s":
		c = &cipher.Substitution{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	case "t":
		c = &cipher.Transposition{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	case "v":
		c = &cipher.Vigenere{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	default:
		panic("Unknown crypto system")

	}

	switch operation {
	case "e":
		encrypted, err := c.Encrypt(string(in))
		if err != nil {
			panic(err)
		}
		out, err := os.Create("encrypt.txt")
		if err != nil {
			panic(err)
		}
		_, err = out.WriteString(encrypted)
		if err != nil {
			panic(err)
		}

		fmt.Println("encrypt success. Saved text to 'encrypt.txt'")

	case "d":
		decrypted, err := c.Decrypt(string(in))
		if err != nil {
			panic(err)
		}
		out, err := os.Create("decrypt.txt")
		if err != nil {
			panic(err)
		}
		_, err = out.WriteString(decrypted)
		if err != nil {
			panic(err)
		}

		fmt.Println("decrypt success. Saved text to 'decrypt.txt'")
	default:
		panic("Unknown operation")
	}

	/*fmt.Println("Enter alphabet filename (type '-' for default alphabet): ")
	_, err := fmt.Scan(&alphabetFileName)
	if err != nil {
		panic(err)
	}
	if alphabetFileName == "-" {
		alphabet = cipher.DefaultAlphabet()
	} else {
		alphabet, err = cipher.NewAlphabetFromFile(alphabetFileName)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Enter input filename: ")
	_, err = fmt.Scan(&inputFileName)
	if err != nil {
		panic(err)
	}
	in, _ := os.ReadFile(inputFileName)

	fmt.Println("Enter key file name: ")
	_, err = fmt.Scan(&keyFileName)
	if err != nil {
		panic(err)
	}
	key, _ := os.ReadFile(keyFileName)

	fmt.Println(`Enter requested cryptoSystem ('a' - affine, 'c' - caesar, 'h' - hill, 's' - substitution, 't' - transposition, 'v' - vigenere): `)
	_, err = fmt.Scan(&operation)
	if err != nil {
		panic(err)
	}
	switch operation {
	case "a":
		c = &cipher.Affine{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	case "c":
		c = &cipher.Caesar{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	case "h":
		c = &cipher.Hill{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	case "s":
		c = &cipher.Substitution{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	case "t":
		c = &cipher.Transposition{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	case "v":
		c = &cipher.Vigenere{
			Alphabet: &alphabet,
			Key:      string(key),
		}

	default:
		panic("Unknown crypto system")

	}

	fmt.Println("Enter requested operation ('e' to encrypt, 'd' to decrypt): ")
	_, err = fmt.Scan(&operation)
	if err != nil {
		panic(err)
	}
	switch operation {
	case "e":
		encrypted, err := c.Encrypt(string(in))
		if err != nil {
			panic(err)
		}
		out, err := os.Create("encrypt.txt")
		if err != nil {
			panic(err)
		}
		_, err = out.WriteString(encrypted)
		if err != nil {
			panic(err)
		}

		fmt.Println("encrypt success. Saved text to 'encrypt.txt'")

	case "d":
		decrypted, err := c.Decrypt(string(in))
		if err != nil {
			panic(err)
		}
		out, err := os.Create("decrypt.txt")
		if err != nil {
			panic(err)
		}
		_, err = out.WriteString(decrypted)
		if err != nil {
			panic(err)
		}

		fmt.Println("decrypt success. Saved text to 'decrypt.txt'")
	default:
		panic("Unknown operation")
	}*/

}
