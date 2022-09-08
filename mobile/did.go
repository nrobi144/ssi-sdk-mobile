package mobile

import (
	gocrypto "crypto"

	"github.com/TBD54566975/ssi-sdk/did"
	"github.com/TBD54566975/ssi-sdk/util"
)
type DIDKeyWrapper struct {
	PrivateKey *gocrypto.PrivateKey // TODO update PrivateKey type
	DidKey     string
}

// GenerateDIDKey takes in a key type value that this library supports and constructs a conformant did:key identifier.
// The function returns the associated private key value cast to the generic golang crypto.PrivateKey interface.
// To use the private key, it is recommended to re-cast to the associated type. For example, called with the input
// for a secp256k1 key:
// privKey, didKey, err := GenerateDIDKey(Secp256k1)
// if err != nil { ... }
// // where secp is an import alias to the secp256k1 library we use "github.com/decred/dcrd/dcrec/secp256k1/v4"
// secpPrivKey, ok := privKey.(secp.PrivateKey)
// if !ok { ... }
func GenerateDidKey(kt string) (*DIDKeyWrapper, error) {
	privateKey, didKey, err := did.GenerateDIDKey(StringToKeyType(kt))
	return &DIDKeyWrapper{
		PrivateKey: &privateKey,
		DidKey:     string(*didKey),
	}, err
}

// CreateDIDKey constructs a did:key from a specific key type and its corresponding public key
// This method does not attempt to validate that the provided public key is of the specified key type.
// A safer method is `GenerateDIDKey` which handles key generation based on the provided key type.
func CreateDIDKey(kt string, publicKey []byte) (string, error) {
	didKey, error := did.CreateDIDKey(StringToKeyType(kt), publicKey)
	return string(*didKey), error
}

type DecodedDIDKey struct {
	Data    []byte
	KeyType string
}

// Decode takes a did:key and returns the underlying public key value as bytes, the LD key type, and a possible error
func DecodeDIDKey(d string) (*DecodedDIDKey, error) {
	data, keyType, error := did.DIDKey(d).Decode()
	return &DecodedDIDKey{
		Data:    data,
		KeyType: string(keyType),
	}, error
}

// Expand turns the DID key into a complaint DID Document
func ExpandDIDKey(d string) ([]byte, error) {
	didDoc, err := did.DIDKey(d).Expand()
	if err != nil {
		return util.PrettyJSON(didDoc)
	}
	return []byte{}, err
}
