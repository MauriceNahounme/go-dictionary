package dictionary

import (
	"fmt"
	"sort"
)

type Dictionary map[string]string

func New() Dictionary {
	d := make(map[string]string)
	return d
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Get(word string) (string, bool) {
	definition, found := d[word]
	return definition, found
}

func (d Dictionary) Remove(word string) {
	delete(d, word)
}

func (d Dictionary) List() []string {
	var words []string
	for word := range d {
		words = append(words, word)
	}
	sort.Strings(words)

	var result []string
	for _, word := range words {
		result = append(result, fmt.Sprintf("%s: %s", word, d[word]))
	}
	return result
}
