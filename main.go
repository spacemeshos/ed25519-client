package main

import (
	"bytes"
	"github.com/spacemeshos/ed25519"
)

type zeroReader struct{}
func (zeroReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 0
	}
	return len(buf), nil
}

func main () {
	var zero zeroReader

	public, private, _ := ed25519.GenerateKey(zero)
	message := []byte("test message")

	// sign the message
	sig := ed25519.Sign2(private, message)

	// extract public key from signature and the message
	public1, err := ed25519.ExtractPublicKey(message, sig)

	if err != nil {
		panic("Failed to extract public key from signature and message")
	}

	if bytes.Compare(public, public1) != 0 {
		panic("extracted pub key is incorrect")
	}

	// You need to use Verify2() to verify signatures generated with Sign2()
	if !ed25519.Verify2(public1, message, sig) {
		panic ("failed to verify message signed with sign2")
	}

	println("All is groovey")
}