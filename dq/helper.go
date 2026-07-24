package dq

import (
	"encoding/json"
	"fmt"
)

func DebugEnd(debug bool, prefix string) {
	if debug {
		fmt.Println(prefix + ":End")
	}
}
func DebugStart(debug bool, prefix string) {
	if debug {
		fmt.Println(prefix + ":Start")
	}
}
func DebugStruct(debug bool, prefix string, s any) {
	if debug {
		if prefix != "" {
			fmt.Println(prefix + ":")
		}
		if b, err := json.MarshalIndent(s, "", "  "); err == nil {
			fmt.Println(string(b))
		} else {
			fmt.Println(err.Error())
		}
	}
}
