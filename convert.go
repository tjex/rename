package rnm

import (
	"strings"
)

// TODO: Support Regexp.
func convert(fileName string, opts convertOption) string {
	if !opts.AsRegexp {
		return strings.Replace(fileName, opts.From, opts.To, -1)
	}
	return fileName
}
