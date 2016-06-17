package classpath

import (
	"fmt"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	paths := strings.Split(pathList, pathListSeparator)
	compositeEntry := make([]Entry, 0, len(paths))
	for _, path := range paths {
		compositeEntry = append(compositeEntry, newEntry(path))
	}
	return compositeEntry
}

func (e CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range e {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, fmt.Errorf("class not found: %s", className)
}

func (e CompositeEntry) String() string {
	strs := make([]string, len(e))

	for i, entry := range e {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
