package utilities

import "regexp"

var IsValidStringData = regexp.MustCompile(`^[a-zA-Z0-9 ]{0,280}$`).MatchString
