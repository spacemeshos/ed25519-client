package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/spacemeshos/ed25519"
	"io/ioutil"
	"os"
)

// Signature verification
type SignedMessage struct {
	Text string
	Signature string
	PublicKey string
}

func main() {

	// playground()

	if len(os.Args) < 2 {
		fmt.Println("No provided json file name")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		return
	}

	var sm SignedMessage
	json.Unmarshal([]byte(data), &sm)
	fmt.Printf("Message: %s, Signature: %s\n", sm.Text, sm.Signature)
	msg := []byte(sm.Text)
	sig, _ := hex.DecodeString(sm.Signature[2:])
	pub, _ := ed25519.ExtractPublicKey(msg, sig)
	account2 := pub[12:]

	// verify the signature
	verified := ed25519.Verify2(pub, msg, sig)
	if verified == true {
		fmt.Printf("Valid signature. Message: %s. Account: 0x%x\n", msg, account2 )
	} else {
		fmt.Printf("Invalid signature. Message: %s. Account: 0x%x\n", msg, account2 )
	}

}
