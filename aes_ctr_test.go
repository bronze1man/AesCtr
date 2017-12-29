package AesCtr

import (
	"crypto/aes"
	"crypto/cipher"
	"testing"
)

// If we test exactly 1K blocks, we would generate exact multiples of
// the cipher's block size, and the cipher stream fragments would
// always be wordsize aligned, whereas non-aligned is a more typical
// use-case.
const almost1K = 1024 - 5

// copy from /usr/local/go/src/crypto/cipher/benchmark_test.go:69
func BenchmarkAESCTR1K_goAesCtr(b *testing.B) {
	buf := make([]byte, almost1K)
	b.SetBytes(int64(len(buf)))

	var key [16]byte
	var iv [16]byte
	aes, _ := aes.NewCipher(key[:])
	ctr := cipher.NewCTR(aes, iv[:])

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctr.XORKeyStream(buf, buf)
	}
}

func BenchmarkAESCTR1K_thisAesCtr(b *testing.B) {
	buf := make([]byte, almost1K)
	b.SetBytes(int64(len(buf)))

	var key [16]byte
	var iv [16]byte
	aes, _ := NewCipher(key[:])
	ctr := cipher.NewCTR(aes, iv[:])

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctr.XORKeyStream(buf, buf)
	}
}