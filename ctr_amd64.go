package AesCtr

import "crypto/cipher"
// xorBytes xors the contents of a and b and places the resulting values into
// dst. If a and b are not the same length then the number of bytes processed
// will be equal to the length of shorter of the two. Returns the number
// of bytes processed.
//go:noescape
func xorBytes(dst, a, b []byte) int
//go:noescape
func fillEightBlocks(nr int, xk *uint32, dst, counter *byte)
// streamBufferSize is the number of bytes of encrypted counter values to cache.
const streamBufferSize = 32 * 16
type aesctr struct {
	block   *aesCipherAsm // block cipher
	nr      int
	ctr     [BlockSize]byte        // next value of the counter (big endian)
	buffer  []byte                 // buffer for the encrypted counter values
	storage [streamBufferSize]byte // array backing buffer slice
}
// Assert that aesctr implements the ctrAble interface.

// NewCTR returns a Stream which encrypts/decrypts using the AES block
// cipher in counter mode. The length of iv must be the same as BlockSize.
func (c *aesCipherAsm) NewCTR(iv []byte) cipher.Stream {
	if len(iv) != BlockSize {
		panic("cipher.NewCTR: IV length must equal block size")
	}
	var ac aesctr
	ac.block = c
	ac.nr = len(c.enc)/4 - 1
	copy(ac.ctr[:], iv)
	ac.buffer = ac.storage[:0]
	return &ac
}
func (c *aesctr) refill() {
	// Fill up the buffer with incrementing counters encrypted.
	c.buffer = c.storage[:streamBufferSize]
	for i := 0; i < len(c.buffer); i += BlockSize * 8 {
		fillEightBlocks(c.nr, &c.block.enc[0], &c.buffer[i], &c.ctr[0])
	}
}
func (c *aesctr) XORKeyStream(dst, src []byte) {
	if len(src) > 0 {
		// Assert len(dst) >= len(src)
		_ = dst[len(src)-1]
	}
	for len(src) > 0 {
		if len(c.buffer) == 0 {
			c.refill()
		}
		n := xorBytes(dst, src, c.buffer)
		c.buffer = c.buffer[n:]
		src = src[n:]
		dst = dst[n:]
	}
}