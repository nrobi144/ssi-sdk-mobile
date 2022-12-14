package mobile

import (
	gocrypto "crypto"

	ssi "github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/TBD54566975/ssi-sdk/did"
)

type DIDKeyWrapper struct {
	PrivateKey *gocrypto.PrivateKey // TODO update PrivateKey type
	DIDKey     string
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
func GenerateDIDKey(kt string) (*DIDKeyWrapper, error) {
	privateKey, didKey, err := did.GenerateDIDKey(stringToKeyType(kt))
	return &DIDKeyWrapper{
		PrivateKey: &privateKey,
		DIDKey:     string(*didKey),
	}, err
}

// CreateDIDKey constructs a did:key from a specific key type and its corresponding public key
// This method does not attempt to validate that the provided public key is of the specified key type.
// A safer method is `GenerateDIDKey` which handles key generation based on the provided key type.
func CreateDIDKey(kt string, publicKey []byte) (string, error) {
	didKey, err := did.CreateDIDKey(stringToKeyType(kt), publicKey)
	return string(*didKey), err
}

type DecodedDIDKey struct {
	Data    []byte
	KeyType string
}

// DecodeDIDKey Decode takes a did:key and returns the underlying public key value as bytes, the LD key type, and a possible error
func DecodeDIDKey(d string) (*DecodedDIDKey, error) {
	data, keyType, err := did.DIDKey(d).Decode()
	return &DecodedDIDKey{
		Data:    data,
		KeyType: string(keyType),
	}, err
}

type DIDDocumentMobile struct {
	Controller           string                  `json:"controller,omitempty"`
	AlsoKnownAs          string                  `json:"alsoKnownAs,omitempty"`
	VerificationMethod   *VerificationMethodArray    `json:"verificationMethod,omitempty" validate:"dive"`
	Authentication       *VerificationMethodSetArray `json:"authentication,omitempty" validate:"dive"`
	AssertionMethod      *VerificationMethodSetArray `json:"assertionMethod,omitempty" validate:"dive"`
	KeyAgreement         *VerificationMethodSetArray `json:"keyAgreement,omitempty" validate:"dive"`
	CapabilityInvocation *VerificationMethodSetArray `json:"capabilityInvocation,omitempty" validate:"dive"`
	CapabilityDelegation *VerificationMethodSetArray `json:"capabilityDelegation,omitempty" validate:"dive"`
	// TODO() Services             []Service               `json:"service,omitempty" validate:"dive"`
}

// ExpandDIDKey Expand turns the DID key into a compliant DID Document
func ExpandDIDKey(d string) (*DIDDocumentMobile, error) {
	didDoc, err := did.DIDKey(d).Expand()
	if err != nil {
		return &DIDDocumentMobile{
			Controller: didDoc.Controller,
			AlsoKnownAs: didDoc.AlsoKnownAs,
			VerificationMethod: &VerificationMethodArray{items: didDoc.VerificationMethod},
			Authentication: &VerificationMethodSetArray{items: didDoc.Authentication},
			AssertionMethod: &VerificationMethodSetArray{items: didDoc.AssertionMethod},
			KeyAgreement: &VerificationMethodSetArray{items: didDoc.KeyAgreement},
			CapabilityInvocation: &VerificationMethodSetArray{items: didDoc.CapabilityInvocation},
			CapabilityDelegation: &VerificationMethodSetArray{items: didDoc.CapabilityDelegation},
		}, err
	}
	return nil, err
}

func stringToKeyType(s string) ssi.KeyType {
	return ssi.KeyType(s)
}
