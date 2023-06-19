package utilities

import "regexp"

var IsValidStringData = regexp.MustCompile(`^[A-Za-z0-9!@#$%^&*. _-]{0,280}$`).MatchString
