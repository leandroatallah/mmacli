package cliflags

import (
	"flag"
	"fmt"
)

// Flags
type flagStruct struct {
	name         string
	short        string
	defaultValue bool
	description  string
}

var (
	flagList = []flagStruct{
		{"quick", "q", false, "Skip delays"},
		{"noenter", "n", false, "Skip press enter command"},
		{"mock", "m", false, "Use mock to fill players name"},
		{"preserve", "p", false, "Don't clear screen"},
	}
	CliFlag = make(map[string]*bool)
)

func init() {
	for _, f := range flagList {
		CliFlag[f.name] = flag.Bool(f.name, f.defaultValue, f.description)
		CliFlag[f.name] = flag.Bool(f.short, f.defaultValue, fmt.Sprintf("Alias for %s", f.name))
	}
}

func GetCliFlag(name string) bool {
	if flag, exists := CliFlag[name]; exists {
		return *flag
	}
	return false
}
