package utilities

import "regexp"

var IsValidStringData = regexp.MustCompile(`^[\pL\pM\pZ\pS\pN\pP]{0,280}$`).MatchString
