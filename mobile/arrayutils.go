package mobile

import(
	"github.com/TBD54566975/ssi-sdk/did"
)

type StringCollection interface {
	Add(s string) StringCollection
	Get(i int) string
	Size() int
}

// TODO solve this with generics
type StringArray struct {
	items []string
}

func (array StringArray) Add(s string) StringArray {
	array.items = append(array.items, s)
	return array
}

func (array StringArray) Get(i int) string {
	return array.items[i]
}

func (array StringArray) Size() int {
	return len(array.items)
}

type VerificationMethodArray struct {
	items []did.VerificationMethod
}

func (array VerificationMethodArray) Add(item *did.VerificationMethod) VerificationMethodArray {
	array.items = append(array.items, *item)
	return array
}

func (array VerificationMethodArray) Get(i int) *did.VerificationMethod {
	return &array.items[i]
}

func (array VerificationMethodArray) Size() int {
	return len(array.items)
}

type VerificationMethodSetArray struct {
	items []did.VerificationMethodSet
}

func (array VerificationMethodSetArray) Add(item *did.VerificationMethodSet) VerificationMethodSetArray {
	array.items = append(array.items, *item)
	return array
}

func (array VerificationMethodSetArray) Get(i int) *did.VerificationMethodSet {
	return &array.items[i]
}

func (array VerificationMethodSetArray) Size() int {
	return len(array.items)
}

