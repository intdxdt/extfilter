package extfilter

import (
	"path/filepath"
	"strings"
)

type ExtensionFilter struct {
	filters []string
	dict    map[string]bool
	strict  bool
}

func NewExtensionFilters(filters []string, strict bool) *ExtensionFilter {
	var ext = &ExtensionFilter{
		filters: filters,
		strict:  strict,
	}
	return ext.updateDictionary()
}

// Match on file name
func (ext *ExtensionFilter) Match(fileName string) bool {
	var key = strings.TrimPrefix(filepath.Ext(fileName), ".")
	return ext.dict[key]
}

func (ext *ExtensionFilter) updateDictionary() *ExtensionFilter {
	ext.dict = make(map[string]bool, len(ext.filters)*3)
	for _, k := range ext.filters {
		for _, key := range ext.genKeys(k) {
			ext.dict[key] = true
		}
	}
	return ext
}

// generates filter keys
func (ext *ExtensionFilter) genKeys(key string) []string {
	if ext.strict {
		return []string{key}
	}
	return []string{key, strings.ToUpper(key), strings.ToLower(key)}
}
