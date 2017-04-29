package validators

import (
	"github.com/markbates/validate"
	"fmt"
	"unicode/utf8"
)

type StringLengthInRange struct {
	Name string
	Field string
	Min int
	Max int
	Message string
}

// IsValid checks that string in range of min:max
// if max not present or it equal to 0 it will be equal to string length
func (v *StringLengthInRange) IsValid(errors *validate.Errors) {
	strLength := utf8.RuneCountInString(v.Field)
	if v.Max == 0 {
		v.Max = strLength
	}
	if v.Message == "" {
		v.Message = fmt.Sprintf("String not in range(%s, %s)", v.Min, v.Max)
	}
	if !(strLength >= v.Min && strLength <= v.Max) {
		errors.Add(v.Name, v.Message)
	}
}
