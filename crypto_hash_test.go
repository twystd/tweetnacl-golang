package tweetnacl

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"math/rand"
	"testing"
)

// --- CryptoHash ---

// Adapted from tests/core1.c)
func TestCryptoHash(t *testing.T) {
	message := []byte("testing\n")

	expected := []byte{
		0x24, 0xf9, 0x50, 0xaa,
		0xc7, 0xb9, 0xea, 0x9b,
		0x3c, 0xb7, 0x28, 0x22,
		0x8a, 0x0c, 0x82, 0xb6,
		0x7c, 0x39, 0xe9, 0x6b,
		0x4b, 0x34, 0x47, 0x98,
		0x87, 0x0d, 0x5d, 0xae,
		0xe9, 0x3e, 0x3a, 0xe5,
		0x93, 0x1b, 0xaa, 0xe8,
		0xc7, 0xca, 0xcf, 0xea,
		0x4b, 0x62, 0x94, 0x52,
		0xc3, 0x80, 0x26, 0xa8,
		0x1d, 0x13, 0x8b, 0xc7,
		0xaa, 0xd1, 0xaf, 0x3e,
		0xf7, 0xbf, 0xd5, 0xec,
		0x64, 0x6d, 0x6c, 0x28}

	hash, err := CryptoHash(message)

	if err != nil {
		t.Errorf("crypto_hash: %v", err)
		return
	}

	if hash == nil {
		t.Errorf("crypto_hash: nil")
		return
	}

	if !bytes.Equal(hash, expected) {
		t.Errorf("crypto_hash: invalid SHA-512 hash (%v)", hash)
		return
	}
}

func TestCryptoHashLoop(t *testing.T) {
	for mlen := 0; mlen < ROUNDS; mlen++ {
		message := make([]byte, mlen)

		rand.Read(message)

		reference := sha512.Sum512(message)
		hash, err := CryptoHash(message)

		if err != nil {
			t.Errorf("LOOP TEST: crypto_hash error (%v)", err)
			return
		}

		if !bytes.Equal(hash, reference[:]) {
			t.Errorf("LOOP TEST: invalid hash [%x][%x]", reference, hash)
			return
		}
	}
}

func BenchmarkCryptoHash(b *testing.B) {
	message := []byte("testing\n")

	for i := 0; i < b.N; i++ {
		CryptoHash(message)
	}
}

// --- EXAMPLES ---

func ExampleCryptoHash() {
	message := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor")

	hash, err := CryptoHash(message)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("[%x]", hash)

	// Output: [5dfaeb09829a546d8adcef4437957814b7b2f44a128600ab0e4f5322c6150cf5c33957f13055b9266e370c199bb764d4f38bb277b5f345e890d2e0bb3992c4dd]
}