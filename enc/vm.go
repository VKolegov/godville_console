package enc

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

/*
	this.Vm = function (e) {
		return s(4) + a[e] + s(5);
	}
*/
func Vm(hash string) string {

	var (
		sb strings.Builder
	)

	sb.Grow(25)

	sb.Write(randomString(4))
	sb.Write([]byte(HTable[hash]))
	sb.Write(randomString(5))

	return sb.String()
}

/*
	s = function (e) {
		for (var t = ""; e > t.length;)
			t += i();
		return t;
	};
*/

func randomString(length int) []byte {

	var (
		rString = make([]byte, length)
	)

	for i := 0; i < length; i++ {
		c := randomChar()

		rString[i] = c
	}

	return rString
}

/*
	i = function () {
		var e = Math.floor(62 * Math.random());
		return 10 > e ? e : 36 > e ? String.fromCharCode(e + 55) : String.fromCharCode(e + 61);
	},
*/

// TODO: improve, optimize
func randomChar() byte {
	e := byte(
		math.Floor(
			62.0 * rand.Float64(),
		),
	)

	if 10 > e {
		return strconv.Itoa(int(e))[0]
	} else if 36 > e {
		return e + 55
	} else {
		return e + 61
	}

}
