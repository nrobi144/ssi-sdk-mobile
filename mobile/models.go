package mobile

import (
	ssi "github.com/TBD54566975/ssi-sdk/crypto"
)

func keyTypeToString(kt ssi.KeyType) string {
	return string(kt)
}

func signatureToString(s ssi.SignatureAlgorithm) string {
	return string(s)
}
