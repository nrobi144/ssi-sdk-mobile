package mobile

import (
	ssi "github.com/TBD54566975/ssi-sdk/crypto"
)

func keyTypeToString(kt ssi.KeyType) string {
	return string(kt)
}

func StringToKeyType(s string) ssi.KeyType {
	return ssi.KeyType(s)
}

func signatureToString(s ssi.SignatureAlgorithm)  string {
	return string(s)
}