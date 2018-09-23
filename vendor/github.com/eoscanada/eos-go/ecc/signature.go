package ecc

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go/btcsuite/btcd/btcec"
	"github.com/eoscanada/eos-go/btcsuite/btcutil/base58"
)

// Signature represents a signature for some hash
type Signature []byte

// Verify checks the signature against the pubKey. `hash` is a sha256
// hash of the payload to verify.
func (s Signature) Verify(hash []byte, pubKey PublicKey) bool {
	recoveredKey, _, err := btcec.RecoverCompact(btcec.S256(), s, hash)
	if err != nil {
		return false
	}
	key, err := pubKey.Key()
	if err != nil {
		return false
	}
	if recoveredKey.IsEqual(key) {
		return true
	}
	return false
}

// PublicKey retrieves the public key, but requires the
// payload.. that's the way to validate the signature. Use Verify() if
// you only want to validate.
func (s Signature) PublicKey(hash []byte) (PublicKey, error) {
	recoveredKey, _, err := btcec.RecoverCompact(btcec.S256(), s, hash)
	if err != nil {
		return nil, err
	}

	return PublicKey(recoveredKey.SerializeCompressed()), nil
}

func (s Signature) String() string {
	checksum := ripemd160checksum(s)
	buf := append(s[:], checksum...)
	return "EOS" + base58.Encode(buf)
}

func NewSignature(fromText string) (Signature, error) {
	sigbytes := base58.Decode(fromText[3:]) // simply remove the `EOS` in front..

	content := sigbytes[:len(sigbytes)-4]
	checksum := sigbytes[len(sigbytes)-4:]
	verifyChecksum := ripemd160checksum(content)
	if !bytes.Equal(verifyChecksum, checksum) {
		return nil, fmt.Errorf("signature checksum failed, found %x expected %x", verifyChecksum, checksum)
	}

	return Signature(content), nil
}

func (a Signature) MarshalBinary() ([]byte, error) {
	return append(bytes.Repeat([]byte{0}, 66-len(a)), a...), nil
}

func (a *Signature) UnmarshalBinary(data []byte) error {
	*a = Signature(data)
	return nil
}

func (a *Signature) UnmarshalBinarySize() int { return 66 }

func (a Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a *Signature) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	*a, err = NewSignature(s)

	return
}
