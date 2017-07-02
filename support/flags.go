package support

import (
	"fmt"
)

// See https://lawlessguy.wordpress.com/2013/07/23/filling-a-slice-using-command-line-flags-in-go-golang/

// StringParamsSet describes multiple parameters set of strings
type StringParamsSet []string

// String my be called by package
func (s *StringParamsSet) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set sets the value
func (s *StringParamsSet) Set(value string) error {
	*s = append(*s, value)
	return nil
}
