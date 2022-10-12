package types

import (
	"encoding/json"
	"strings"
)

type Status string

const (
	yes Status = "YES"
	no  Status = "NO"
)

// UnmarshalJSON This custom Unmarshal method is written to validate enum status field.
func (s *Status) UnmarshalJSON(b []byte) error {
	var temp string

	err := json.Unmarshal(b, &temp)
	if err != nil {
		return &json.UnmarshalTypeError{Field: "status"}
	}

	temp = strings.TrimSpace(temp)

	switch temp {
	case "YES":
		*s = yes
	case "NO":
		*s = no
	case "":
		*s = ""
	default:
		return &json.UnmarshalTypeError{Field: "status"}
	}

	return nil
}
