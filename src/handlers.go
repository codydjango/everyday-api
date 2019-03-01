package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
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

// GetOrdinal returns ordinals for integers
func GetOrdinal(i int) string {
	j := i % 10
	k := i % 100

	if j == 1 && k != 11 {
		return "st"
	}

	if j == 2 && k != 12 {
		return "nd"
	}

	if j == 3 && k != 13 {
		return "rd"
	}

	return "th"
}

// AddressNonceLookup keep track of the nonce for each address
var AddressNonceLookup = make(map[string]int)

// Claim is a little struct for helping with authentication
type Claim struct {
	Signature     string `json:"signature,omitempty"`
	PublicAddress string `json:"publicAddress,omitempty"`
	Verified      bool   `json:"verified"`
	Challenge     string `json:"challenge"`
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

// VersionHandler controller returns version of the API
func VersionHandler(responseWriter http.ResponseWriter, request *http.Request) {
	version := "0.0.1"

	keyMapping := make(map[string]string)
	keyMapping["version"] = version

	json, err := json.Marshal(keyMapping)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(json)
}

// AddressNonceHandler controller
func AddressNonceHandler(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	publicKey := vars["address"]

	if AddressNonceLookup[publicKey] == 0 {
		AddressNonceLookup[publicKey]++
	}

	keyMapping := make(map[string]string)
	keyMapping["nonce"] = fmt.Sprintf("%d", AddressNonceLookup[publicKey])
	keyMapping["publicKey"] = publicKey

	json, err := json.Marshal(keyMapping)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(json)
}

// DecodeClaim is the function that decodes the json into a claim struct
func DecodeClaim(r io.ReadCloser) (x *Claim, err error) {
	// x = new(Claim)
	x = &Claim{}
	err = json.NewDecoder(r).Decode(x)
	return
}

// AuthenticationHandler controller
func AuthenticationHandler(responseWriter http.ResponseWriter, request *http.Request) {
	log.Println("authentication handler")

	claim, err := DecodeClaim(request.Body)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	claim.updateChallenge()
	claim.verify()

	if claim.Verified == true {
		claim.updateNonce()
	}

	json, err := json.Marshal(claim)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(json)
}
