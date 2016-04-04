package tweetnacl

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
)

var CIPHERTEXT = []byte{
	0xf3, 0xff, 0xc7, 0x70,
	0x3f, 0x94, 0x00, 0xe5,
	0x2a, 0x7d, 0xfb, 0x4b,
	0x3d, 0x33, 0x05, 0xd9,
	0x8e, 0x99, 0x3b, 0x9f,
	0x48, 0x68, 0x12, 0x73,
	0xc2, 0x96, 0x50, 0xba,
	0x32, 0xfc, 0x76, 0xce,
	0x48, 0x33, 0x2e, 0xa7,
	0x16, 0x4d, 0x96, 0xa4,
	0x47, 0x6f, 0xb8, 0xc5,
	0x31, 0xa1, 0x18, 0x6a,
	0xc0, 0xdf, 0xc1, 0x7c,
	0x98, 0xdc, 0xe8, 0x7b,
	0x4d, 0xa7, 0xf0, 0x11,
	0xec, 0x48, 0xc9, 0x72,
	0x71, 0xd2, 0xc2, 0x0f,
	0x9b, 0x92, 0x8f, 0xe2,
	0x27, 0x0d, 0x6f, 0xb8,
	0x63, 0xd5, 0x17, 0x38,
	0xb4, 0x8e, 0xee, 0xe3,
	0x14, 0xa7, 0xcc, 0x8a,
	0xb9, 0x32, 0x16, 0x45,
	0x48, 0xe5, 0x26, 0xae,
	0x90, 0x22, 0x43, 0x68,
	0x51, 0x7a, 0xcf, 0xea,
	0xbd, 0x6b, 0xb3, 0x73,
	0x2b, 0xc0, 0xe9, 0xda,
	0x99, 0x83, 0x2b, 0x61,
	0xca, 0x01, 0xb6, 0xde,
	0x56, 0x24, 0x4a, 0x9e,
	0x88, 0xd5, 0xf9, 0xb3,
	0x79, 0x73, 0xf6, 0x22,
	0xa4, 0x3d, 0x14, 0xa6,
	0x59, 0x9b, 0x1f, 0x65,
	0x4c, 0xb4, 0x5a, 0x74,
	0xe3, 0x55, 0xa5}

var MESSAGE = []byte{
	0xbe, 0x07, 0x5f, 0xc5,
	0x3c, 0x81, 0xf2, 0xd5,
	0xcf, 0x14, 0x13, 0x16,
	0xeb, 0xeb, 0x0c, 0x7b,
	0x52, 0x28, 0xc5, 0x2a,
	0x4c, 0x62, 0xcb, 0xd4,
	0x4b, 0x66, 0x84, 0x9b,
	0x64, 0x24, 0x4f, 0xfc,
	0xe5, 0xec, 0xba, 0xaf,
	0x33, 0xbd, 0x75, 0x1a,
	0x1a, 0xc7, 0x28, 0xd4,
	0x5e, 0x6c, 0x61, 0x29,
	0x6c, 0xdc, 0x3c, 0x01,
	0x23, 0x35, 0x61, 0xf4,
	0x1d, 0xb6, 0x6c, 0xce,
	0x31, 0x4a, 0xdb, 0x31,
	0x0e, 0x3b, 0xe8, 0x25,
	0x0c, 0x46, 0xf0, 0x6d,
	0xce, 0xea, 0x3a, 0x7f,
	0xa1, 0x34, 0x80, 0x57,
	0xe2, 0xf6, 0x55, 0x6a,
	0xd6, 0xb1, 0x31, 0x8a,
	0x02, 0x4a, 0x83, 0x8f,
	0x21, 0xaf, 0x1f, 0xde,
	0x04, 0x89, 0x77, 0xeb,
	0x48, 0xf5, 0x9f, 0xfd,
	0x49, 0x24, 0xca, 0x1c,
	0x60, 0x90, 0x2e, 0x52,
	0xf0, 0xa0, 0x89, 0xbc,
	0x76, 0x89, 0x70, 0x40,
	0xe0, 0x82, 0xf9, 0x37,
	0x76, 0x38, 0x48, 0x64,
	0x5e, 0x07, 0x05}

var NONCE = []byte{
	0x69, 0x69, 0x6e, 0xe9,
	0x55, 0xb6, 0x2b, 0x73,
	0xcd, 0x62, 0xbd, 0xa8,
	0x75, 0xfc, 0x73, 0xd6,
	0x82, 0x19, 0xe0, 0x03,
	0x6b, 0x7a, 0x0b, 0x37}

var ALICEPK = []byte{
	0x85, 0x20, 0xf0, 0x09,
	0x89, 0x30, 0xa7, 0x54,
	0x74, 0x8b, 0x7d, 0xdc,
	0xb4, 0x3e, 0xf7, 0x5a,
	0x0d, 0xbf, 0x3a, 0x0d,
	0x26, 0x38, 0x1a, 0xf4,
	0xeb, 0xa4, 0xa9, 0x8e,
	0xaa, 0x9b, 0x4e, 0x6a}

var ALICESK = []byte{
	0x77, 0x07, 0x6d, 0x0a,
	0x73, 0x18, 0xa5, 0x7d,
	0x3c, 0x16, 0xc1, 0x72,
	0x51, 0xb2, 0x66, 0x45,
	0xdf, 0x4c, 0x2f, 0x87,
	0xeb, 0xc0, 0x99, 0x2a,
	0xb1, 0x77, 0xfb, 0xa5,
	0x1d, 0xb9, 0x2c, 0x2a}

var BOBPK = []byte{
	0xde, 0x9e, 0xdb, 0x7d,
	0x7b, 0x7d, 0xc1, 0xb4,
	0xd3, 0x5b, 0x61, 0xc2,
	0xec, 0xe4, 0x35, 0x37,
	0x3f, 0x83, 0x43, 0xc8,
	0x5b, 0x78, 0x67, 0x4d,
	0xad, 0xfc, 0x7e, 0x14,
	0x6f, 0x88, 0x2b, 0x4f}

var BOBSK = []byte{
	0x5d, 0xab, 0x08, 0x7e,
	0x62, 0x4a, 0x8a, 0x4b,
	0x79, 0xe1, 0x7f, 0x8b,
	0x83, 0x80, 0x0e, 0xe6,
	0x6f, 0x3b, 0xb1, 0x29,
	0x26, 0x18, 0xb6, 0xfd,
	0x1c, 0x2f, 0x8b, 0x27,
	0xff, 0x88, 0xe0, 0xeb}

var KEY = []byte{
	0x1b, 0x27, 0x55, 0x64,
	0x73, 0xe9, 0x85, 0xd4,
	0x62, 0xcd, 0x51, 0x19,
	0x7a, 0x9a, 0x46, 0xc7,
	0x60, 0x09, 0x54, 0x9e,
	0xac, 0x64, 0x74, 0xf2,
	0x06, 0xc4, 0xee, 0x08,
	0x44, 0xf6, 0x83, 0x89}

const ROUNDS int = 100 // TODO change to 10000 when done

// --- CryptoBoxKeyPair ---

func TestCryptoBoxKeyPair(t *testing.T) {
	keypair, err := CryptoBoxKeyPair()

	if err != nil {
		t.Errorf("cryptobox_keypair: %v", err)
		return
	}

	if keypair == nil {
		t.Errorf("cryptobox_keypair: nil")
		return
	}

	if keypair.PublicKey == nil || len(keypair.PublicKey) != 32 {
		t.Errorf("cryptobox_keypair: invalid public key")
		return
	}

	if keypair.SecretKey == nil || len(keypair.SecretKey) != 32 {
		t.Errorf("cryptobox_keypair: invalid secret key")
		return
	}
}

func BenchmarkCryptoBoxKeyPair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CryptoBoxKeyPair()
	}
}

// --- CryptoBox ---

// Adapted from tests/box.c
func TestCryptoBox(t *testing.T) {
	ciphertext, err := CryptoBox(MESSAGE, NONCE, BOBPK, ALICESK)

	if err != nil {
		t.Errorf("cryptobox: %v", err)
		return
	}

	if ciphertext == nil {
		t.Errorf("cryptobox: nil")
		return
	}

	if !bytes.Equal(ciphertext, CIPHERTEXT) {
		t.Errorf("cryptobox: invalid ciphertext (%v)", ciphertext)
		return
	}
}

func BenchmarkCryptoBox(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CryptoBox(MESSAGE, NONCE, BOBPK, ALICESK)
	}
}

// --- CryptoBoxOpen ---

// Adapted from tests/box2.c
func TestCryptoBoxOpen(t *testing.T) {
	plaintext, err := CryptoBoxOpen(CIPHERTEXT, NONCE, ALICEPK, BOBSK)

	if err != nil {
		t.Errorf("cryptobox_open: %v", err)
		return
	}

	if plaintext == nil {
		t.Errorf("cryptobox_open: nil")
		return
	}

	if !bytes.Equal(plaintext, MESSAGE) {
		t.Errorf("cryptobox_open: invalid plaintext (%v)", plaintext)
		return
	}
}

func BenchmarkCryptoBoxOpen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CryptoBoxOpen(CIPHERTEXT, NONCE, ALICEPK, BOBSK)
	}
}

// --- CryptoBoxBeforeNM ---

func TestCryptoBoxBeforeNM(t *testing.T) {
	key, err := CryptoBoxBeforeNM(BOBPK, ALICESK)

	if err != nil {
		t.Errorf("cryptobox_beforenm: %v", err)
		return
	}

	if key == nil {
		t.Errorf("cryptobox_beforenm: nil")
		return
	}

	if !bytes.Equal(key, KEY) {
		t.Errorf("cryptobox_beforenm: invalid shared key (%v)", key)
		return
	}
}

func BenchmarkCryptoBoxBeforeNM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CryptoBoxBeforeNM(BOBPK, ALICESK)
	}
}

// --- CryptoBoxAfterNM ---

func TestCryptoBoxAfteNM(t *testing.T) {
	key, err := CryptoBoxBeforeNM(BOBPK, ALICESK)
	ciphertext, err := CryptoBoxAfterNM(MESSAGE, NONCE, key)

	if err != nil {
		t.Errorf("cryptobox_afternm: %v", err)
		return
	}

	if ciphertext == nil {
		t.Errorf("cryptobox_afternm: nil")
		return
	}

	if !bytes.Equal(ciphertext, CIPHERTEXT) {
		t.Errorf("cryptobox_afternm: invalid ciphertext (%v)", ciphertext)
		return
	}
}

func BenchmarkCryptoBoxAfterNM(b *testing.B) {
	key, _ := CryptoBoxBeforeNM(BOBPK, ALICESK)
	for i := 0; i < b.N; i++ {
		CryptoBoxAfterNM(MESSAGE, NONCE, key)
	}
}

// --- LOOP TESTS ---

// Adapted from tests/box7.c
func TestCryptoBoxLoop(t *testing.T) {
	for mlen := 0; mlen < ROUNDS; mlen++ {
		alice, _ := CryptoBoxKeyPair()
		bob, _ := CryptoBoxKeyPair()
		message := make([]byte, mlen)
		nonce := make([]byte, 24)

		rand.Read(nonce)
		rand.Read(message)

		ciphertext, err := CryptoBox(message, nonce, bob.PublicKey, alice.SecretKey)

		if err != nil {
			t.Errorf("LOOP TEST: encryption error (%v)", err)
			return
		}

		plaintext, err := CryptoBoxOpen(ciphertext, nonce, alice.PublicKey, bob.SecretKey)

		if err != nil {
			t.Errorf("LOOP TEST: decryption error (%v)", err)
			return
		}

		if !bytes.Equal(message, plaintext) {
			t.Errorf("LOOP TEST: bad decryption (original :%v)", message)
			t.Errorf("LOOP TEST: bad decryption (decrypted:%v)", plaintext)
			return
		}
	}
}

// Adapted from tests/box8.c
func TestCryptoBoxCorruptedCiphertext(t *testing.T) {
	for mlen := 0; mlen < ROUNDS; mlen++ {
		alice, _ := CryptoBoxKeyPair()
		bob, _ := CryptoBoxKeyPair()
		message := make([]byte, mlen)
		nonce := make([]byte, 24)
		caught := 0

		rand.Read(nonce)
		rand.Read(message)

		ciphertext, err := CryptoBox(message, nonce, bob.PublicKey, alice.SecretKey)

		if err != nil {
			t.Errorf("CORRUPTED CIPHERTEXT: encryption error (%v)", err)
			return
		}

		for caught < 10 {
			ix := rand.Intn(len(ciphertext))
			b := byte(rand.Intn(256))

			ciphertext[ix] = b
			plaintext, err := CryptoBoxOpen(ciphertext, nonce, alice.PublicKey, bob.SecretKey)

			if err == nil {
				if !bytes.Equal(message, plaintext) {
					t.Errorf("Forgery!!!")
					return
				}
			}

			caught++
		}
	}

}

// --- EXAMPLES ---

func ExampleCryptoBox() {
	message := []byte("Neque porro quisquam est qui dolorem ipsum quia dolor sit amet")
	nonce := make([]byte, 24)
	alice, _ := CryptoBoxKeyPair()
	bob, _ := CryptoBoxKeyPair()

	rand.Read(nonce)

	ciphertext, err := CryptoBox(message, nonce, bob.PublicKey, alice.SecretKey)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	plaintext, err := CryptoBoxOpen(ciphertext, nonce, alice.PublicKey, bob.SecretKey)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("[%s]", string(plaintext))

	// Output: [Neque porro quisquam est qui dolorem ipsum quia dolor sit amet]
}

func ExampleCryptoBoxNM() {
	message := []byte("Neque porro quisquam est qui dolorem ipsum quia dolor sit amet")
	nonce := make([]byte, 24)
	alice, _ := CryptoBoxKeyPair()
	bob, _ := CryptoBoxKeyPair()

	rand.Read(nonce)

	key, err := CryptoBoxBeforeNM(bob.PublicKey, alice.SecretKey)
	ciphertext, err := CryptoBoxAfterNM(message, nonce, key)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	plaintext, err := CryptoBoxOpen(ciphertext, nonce, alice.PublicKey, bob.SecretKey)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("[%s]", string(plaintext))

	// Output: [Neque porro quisquam est qui dolorem ipsum quia dolor sit amet]
}