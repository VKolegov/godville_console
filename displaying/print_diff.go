package displaying

import (
	"strconv"
	"strings"
)

func appendDiff(curr, last int, sb *strings.Builder) {

	diff := curr - last

	if diff != 0 {

		diffStr := strconv.Itoa(diff)
		sb.WriteString(" (")

		if diff < 0 {
			sb.WriteString(diffStr)
		} else {
			sb.WriteByte('+')
			sb.WriteString(diffStr)
		}

		sb.WriteByte(')')
	}
}
