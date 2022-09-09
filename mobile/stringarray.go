package mobile

type StringCollection interface {
	Add(s string) StringCollection
	Get(i int) string
	Size() int
}

type StringArray struct {
	items []string
}

func (stringArray StringArray) Add(s string) StringArray {
	stringArray.items = append(stringArray.items, s)
	return stringArray
}

func (stringArray StringArray) Get(i int) string {
	return stringArray.items[i]
}

func (stringArray StringArray) Size() int {
	return len(stringArray.items)
}
