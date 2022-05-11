package cipher

type Operation int

const (
	Encrypt Operation = iota
	Decrypt
)

type Cipher interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
	//validateKey()
}

type Keys struct {
	key rune
	pos int
}

/*type CryptoSystem int

const (
	caesar CryptoSystem = iota
	affine
	simpleSubstitution
	hill
	transposition
	vigenere
)*/
