package encryption

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"log"
)

type CFBCodec struct {
	Reader *cipher.StreamReader
	Writer *cipher.StreamWriter
}

func NewCFBCodec(key *[]byte) *CFBCodec {
	block, _ := aes.NewCipher(*key)
	return &CFBCodec{
		Reader: &cipher.StreamReader{
			S: NewDecrypt(block, *key),
			R: nil,
		},
		Writer: &cipher.StreamWriter{
			S:   NewEncrypt(block, *key),
			W:   nil,
			Err: nil,
		},
	}
}

func (codec *CFBCodec) Encrypt(writer io.Writer) {
	codec.Writer.W = writer
}

func (codec *CFBCodec) Decrypt(reader *bufio.Reader) {
	codec.Reader.R = reader
}

type CFB8 struct {
	b         cipher.Block
	blockSize int
	in        []byte
	out       []byte

	decrypt bool
}

func NewCFB8(block cipher.Block, iv []byte, decrypt bool) cipher.Stream {
	blockSize := block.BlockSize()

	if len(iv) != blockSize {
		log.Panic("IV should have same length as block size: ", len(iv), " vs ", blockSize)
	}

	x := &CFB8{
		b:         block,
		blockSize: blockSize,
		out:       make([]byte, blockSize),
		in:        make([]byte, blockSize),
		decrypt:   decrypt,
	}
	copy(x.in, iv)

	return x
}

func NewEncrypt(block cipher.Block, iv []byte) cipher.Stream {
	return NewCFB8(block, iv, false)
}

func NewDecrypt(block cipher.Block, iv []byte) cipher.Stream {
	return NewCFB8(block, iv, true)
}

func (x *CFB8) XORKeyStream(dst, src []byte) {
	for i := range src {
		x.b.Encrypt(x.out, x.in)
		copy(x.in[:x.blockSize-1], x.in[1:])
		if x.decrypt {
			x.in[x.blockSize-1] = src[i]
		}
		dst[i] = src[i] ^ x.out[0]
		if !x.decrypt {
			x.in[x.blockSize-1] = dst[i]
		}
	}
}
