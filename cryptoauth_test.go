package tweetnacl

import (
	"bytes"
	"math/rand"
	"testing"
)

// --- CryptoOneTimeAuth ---

// Adapted from tests/onetimeauth.c
func TestCryptoOneTimeAuth(t *testing.T) {
	message := []byte{
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

	key := []byte{
		0xee, 0xa6, 0xa7, 0x25,
		0x1c, 0x1e, 0x72, 0x91,
		0x6d, 0x11, 0xc2, 0xcb,
		0x21, 0x4d, 0x3c, 0x25,
		0x25, 0x39, 0x12, 0x1d,
		0x8e, 0x23, 0x4e, 0x65,
		0x2d, 0x65, 0x1f, 0xa4,
		0xc8, 0xcf, 0xf8, 0x80}

	expected := []byte{
		0xf3, 0xff, 0xc7, 0x70,
		0x3f, 0x94, 0x00, 0xe5,
		0x2a, 0x7d, 0xfb, 0x4b,
		0x3d, 0x33, 0x05, 0xd9}

	authenticator, err := CryptoOneTimeAuth(message, key)

	if err != nil {
		t.Errorf("crypto_onetimeauth: %v", err)
		return
	}

	if authenticator == nil {
		t.Errorf("crypto_onetimeauth: nil")
		return
	}

	if !bytes.Equal(authenticator, expected) {
		t.Errorf("crypto_onetimeauth: invalid one-time authenticator ([% x])", authenticator)
		return
	}
}

func BenchmarkCryptoOneTimeAuth(b *testing.B) {
	message := []byte{
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

	key := []byte{
		0xee, 0xa6, 0xa7, 0x25,
		0x1c, 0x1e, 0x72, 0x91,
		0x6d, 0x11, 0xc2, 0xcb,
		0x21, 0x4d, 0x3c, 0x25,
		0x25, 0x39, 0x12, 0x1d,
		0x8e, 0x23, 0x4e, 0x65,
		0x2d, 0x65, 0x1f, 0xa4,
		0xc8, 0xcf, 0xf8, 0x80}

	for i := 0; i < b.N; i++ {
		CryptoOneTimeAuth(message, key)
	}
}

// --- CryptoOneTimeAuthVerify ---

// Adapted from tests/onetimeauth2.c
func TestCryptoOneTimeAuthVerify(t *testing.T) {
	authenticator := []byte{
		0xf3, 0xff, 0xc7, 0x70,
		0x3f, 0x94, 0x00, 0xe5,
		0x2a, 0x7d, 0xfb, 0x4b,
		0x3d, 0x33, 0x05, 0xd9}

	message := []byte{
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

	key := []byte{
		0xee, 0xa6, 0xa7, 0x25,
		0x1c, 0x1e, 0x72, 0x91,
		0x6d, 0x11, 0xc2, 0xcb,
		0x21, 0x4d, 0x3c, 0x25,
		0x25, 0x39, 0x12, 0x1d,
		0x8e, 0x23, 0x4e, 0x65,
		0x2d, 0x65, 0x1f, 0xa4,
		0xc8, 0xcf, 0xf8, 0x80}

	ok, err := CryptoOneTimeAuthVerify(authenticator, message, key)

	if err != nil {
		t.Errorf("crypto_onetimeauth_verify: %v", err)
		return
	}

	if !ok {
		t.Errorf("crypto_onetimeauth_verify: failed")
		return
	}

}

// Adapted from tests/onetimeauth7.c
func TestCryptoOneTimeAuthVerifyWithMutations(t *testing.T) {
	for clen := 0; clen < ROUNDS; clen++ {
		key := make([]byte, ONETIMEAUTH_KEYBYTES)
		message := make([]byte, clen)

		rand.Read(key)
		rand.Read(message)

		authenticator, err := CryptoOneTimeAuth(message, key)

		if err != nil {
			t.Errorf("crypto_onetimeauth: %v", err)
			return
		}

		if len(authenticator) != ONETIMEAUTH_BYTES {
			t.Errorf("crypto_onetimeauth_verify: %v", err)
			return
		}

		if clen > 0 {
			mx := rand.Intn(clen)
			ax := rand.Intn(ONETIMEAUTH_BYTES)
			m := make([]byte, len(message))
			a := make([]byte, len(authenticator))

			copy(m, message)
			copy(a, authenticator)

			m[mx] = message[mx] + byte(1+rand.Intn(255))
			a[ax] = authenticator[ax] + byte(1+rand.Intn(255))

			ok, err := CryptoOneTimeAuthVerify(authenticator, m, key)

			if err == nil {
				t.Errorf("crypto_onetimeauth_verify: %s", "Verified forgery with altered message")
				return
			}
			if ok {
				t.Errorf("crypto_onetimeauth_verify: %s", "Verified forgery with altered message")
				return
			}

			ok, err = CryptoOneTimeAuthVerify(a, message, key)

			if err == nil {
				t.Errorf("crypto_onetimeauth_verify: %s", "Verified forgery with altered authentication")
				return
			}

			if ok {
				t.Errorf("crypto_onetimeauth_verify: %s", "Verified forgery with altered authentication")
				return
			}

			ok, err = CryptoOneTimeAuthVerify(a, m, key)

			if err == nil {
				t.Errorf("crypto_onetimeauth_verify: %s", "Verified forgery with altered message and authentication")
				return
			}
			if ok {
				t.Errorf("crypto_onetimeauth_verify: %s", "Verified forgery with altered message and authentication")
				return
			}
		}
	}
}

func BenchmarkCryptoOneTimeAuthVerify(b *testing.B) {
	authenticator := []byte{
		0xf3, 0xff, 0xc7, 0x70,
		0x3f, 0x94, 0x00, 0xe5,
		0x2a, 0x7d, 0xfb, 0x4b,
		0x3d, 0x33, 0x05, 0xd9}

	message := []byte{
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

	key := []byte{
		0xee, 0xa6, 0xa7, 0x25,
		0x1c, 0x1e, 0x72, 0x91,
		0x6d, 0x11, 0xc2, 0xcb,
		0x21, 0x4d, 0x3c, 0x25,
		0x25, 0x39, 0x12, 0x1d,
		0x8e, 0x23, 0x4e, 0x65,
		0x2d, 0x65, 0x1f, 0xa4,
		0xc8, 0xcf, 0xf8, 0x80}

	for i := 0; i < b.N; i++ {
		CryptoOneTimeAuthVerify(authenticator, message, key)
	}
}
