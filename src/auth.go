/*
Structs and functions for authentication behaviours.
*/

package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// hashChallenge pads the challenge with eth padding and then returns the hash.
// more info here: https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L404
func hashChallenge(challenge string) []byte {
	challengeBytes := []byte(challenge)
	lengthOfChallengeBytes := len(challengeBytes)
	paddedChallenge := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", lengthOfChallengeBytes, challengeBytes)
	return crypto.Keccak256([]byte(paddedChallenge))
}

// verifySignature takes public
func verifySignature(publicKey, signatureHex string, challenge string) bool {
	publicKeyAddress := common.HexToAddress(publicKey)

	// MustDecode decodes a hex string with 0x prefix. Returns a slice of bytes
	signatureBytes := hexutil.MustDecode(signatureHex)
	hashedChallenge := hashChallenge(challenge)

	// I'm still learning why this is the case.
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if signatureBytes[64] != 27 && signatureBytes[64] != 28 {
		return false
	}

	signatureBytes[64] -= 27
	signaturePublicKey, err := crypto.SigToPub(hashedChallenge, signatureBytes)
	if err != nil {
		return false
	}

	return publicKeyAddress == crypto.PubkeyToAddress(*signaturePublicKey)
}

// AddressNonceLookup keep track of the nonce for each address
var AddressNonceLookup = make(map[string]int)

// Claim is a little struct for helping with authentication
type Claim struct {
	Signature     string `json:"signature,omitempty"`
	PublicAddress string `json:"publicAddress,omitempty"`
	Verified      bool   `json:"verified"`
	Challenge     string `json:"-"`
}

func (c Claim) getNonce() int {
	return AddressNonceLookup[c.PublicAddress]
}

func (c *Claim) updateChallenge() {
	values := []interface{}{c.getNonce(), GetOrdinal(c.getNonce())}
	c.Challenge = fmt.Sprintf("I'm signing into my everyday account for the %d%s time", values...)
	return
}

func (c *Claim) newChallenge() {
	AddressNonceLookup[c.PublicAddress]++
	c.updateChallenge()
	return
}

func (c *Claim) verify() {
	c.Verified = verifySignature(c.PublicAddress, c.Signature, c.Challenge)
	return
}

func (c *Claim) updateNonce() {
	AddressNonceLookup[c.PublicAddress]++
}

func (c *Claim) isValid() bool {
	return len(c.Signature) == 132 &&
		c.Signature[0:2] == "0x" &&
		len(c.PublicAddress) == 42 &&
		c.PublicAddress[0:2] == "0x"
}
