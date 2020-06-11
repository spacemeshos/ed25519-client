package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/spacemeshos/ed25519"
)

type zeroReader struct{}

func (zeroReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 0
	}
	return len(buf), nil
}

func main() {
	var zero zeroReader

	public, private, _ := ed25519.GenerateKey(zero)
	message := []byte("hello spacemesh")

	fmt.Printf("Message utf-8 hex: ")
	for i := 0; i < len(message); i++ {
		fmt.Printf("%x", message[i])
	}
	println()

	// sign the message
	sig := ed25519.Sign2(private, message)

	// extract the public key from signature and the message
	public1, err := ed25519.ExtractPublicKey(message, sig)

	if err != nil {
		panic("Failed to extract public key from signature and message")
	}

	if bytes.Compare(public, public1) != 0 {
		panic("extracted pub key is incorrect")
	}

	// You need to use Verify2() to verify signatures generated with Sign2()
	if !ed25519.Verify2(public1, message, sig) {
		panic("failed to verify message signed with sign2")
	}


	// pub, _ := hex.DecodeString("40eebccd54af06be0c7b4df54fcd5295754e6732eb64d5f8340751c72032db8a")
	pri, _ := hex.DecodeString("fe242b00cf1d3d025f9eecc129a479d43259f061ccc3e4167b7678dee2ac53b740eebccd54af06be0c7b4df54fcd5295754e6732eb64d5f8340751c72032db8a")
	sig = ed25519.Sign2(pri, message)
	fmt.Printf("Signature hex: %x \n", sig)
	expectedSig, _ := hex.DecodeString("20122e1b75012f5a816b0db5a0673a1d5599112120a6a6b8616c9768dcc275086dfa556c92424d9d5e9cef181b92307a8176caf8e403491e416f0a7690877b01")

	if bytes.Compare(sig, expectedSig) != 0 {
		panic("Signature mismatch")
	}

	println("All is cool.")
}
