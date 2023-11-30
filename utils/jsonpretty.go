package utils

import (
	"bytes"
	"encoding/json"
)

/*
provides functionality to parse json files and make it more readable.
*/

// pretty print json
func PrettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
