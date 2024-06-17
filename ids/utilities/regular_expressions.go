package utilities

import "regexp"

var IsValidStringID = regexp.MustCompile(`^[A-Za-z0-9_]{0,30}$`).MatchString
