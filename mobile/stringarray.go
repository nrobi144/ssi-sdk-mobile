package mobile

type StringArray interface {
	Add(s string) StringArray
	Get(i int) string
	Size() int
}

type StringArrayImpl struct {
	items []string
}

func (stringArray StringArrayImpl) Add(s string) StringArray {
	stringArray.items = append(stringArray.items, s)
	return stringArray
}

func (stringArray StringArrayImpl) Get(i int) string {
	return stringArray.items[i]
}

func (stringArray StringArrayImpl) Size() int {
	return len(stringArray.items)
}
