// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2018-2021 Sylabs, Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license.  Please
// consult LICENSE.md file distributed with the sources of this project regarding
// your rights to use or distribute this software.

package shell

import "strings"

// ArgsQuoted concatenates a slice of string shell args, quoting each item
func ArgsQuoted(a []string) (quoted string) {
	for _, val := range a {
		quoted = quoted + `"` + Escape(val) + `" `
	}
	quoted = strings.TrimRight(quoted, " ")
	return
}

// Escape performs escaping of shell double quotes, backticks and $ characters.
// Does not escape single quotes - apply EscapeSingleQuotes separately for this.
func Escape(s string) string {
	escaped := strings.Replace(s, `\`, `\\`, -1)
	escaped = strings.Replace(escaped, `"`, `\"`, -1)
	escaped = strings.Replace(escaped, "`", "\\`", -1)
	escaped = strings.Replace(escaped, `$`, `\$`, -1)
	return escaped
}

// EscapeDoubleQuotes performs shell escaping of double quotes only
func EscapeDoubleQuotes(s string) string {
	return strings.Replace(s, `"`, `\"`, -1)
}

// EscapeSingleQuotes performs shell escaping of single quotes only
func EscapeSingleQuotes(s string) string {
	return strings.Replace(s, `'`, `'"'"'`, -1)
}
