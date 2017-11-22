package format

import "strconv"

//FormatFloat will format float64 to the default representation in horizon:
//decimal with exactly 7 digits after the decimal point
func FormatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', 7, 64)
}