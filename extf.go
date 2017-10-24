package extfilter

import (
	"strings"
	"path/filepath"
)

type ExtensionFilter struct {
	filters []string
	extmap  map[string]bool
	strict  bool
}

//new extension filters
func NewExtensionFilters(filters []string, strict bool) *ExtensionFilter {
	ex := &ExtensionFilter{
		filters: filters,
		strict:  strict,
	}
	return ex.updateFilterMap()
}

//Match on file name
func (exf *ExtensionFilter) Match(fileName string) bool {
	key := strings.TrimPrefix(filepath.Ext(fileName), ".")
	return exf.extmap[key]
}

//filter map
func (exf *ExtensionFilter) updateFilterMap() *ExtensionFilter {
	exf.extmap = make(map[string]bool)
	for _, k := range exf.filters {
		for _, key := range exf.genKeys(k) {
			exf.extmap[key] = true
		}
	}
	return exf
}

//generates filter keys
func (exf *ExtensionFilter) genKeys(key string) []string {
	if exf.strict {
		return []string{key}
	}
	return []string{key, strings.ToUpper(key), strings.ToLower(key)}
}
