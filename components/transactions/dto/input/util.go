package input

import "regexp"

var CurrencyRegex = regexp.MustCompile(`^[1-9]([0-9]{1,3})*(?:,?[0-9]{3})*(?:\.[0-9]{1,2})?$`)
