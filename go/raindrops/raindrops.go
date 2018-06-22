package raindrops

import (
	"math"
	"strconv"
)

//Convert - converts numbers to raindrops :mindblown:
func Convert(num int) string {
	rain := ""
	if math.Remainder(float64(num), 3) == 0 {
		rain += "Pling"
	}

	if math.Remainder(float64(num), 5) == 0 {
		rain += "Plang"
	}

	if math.Remainder(float64(num), 7) == 0 {
		rain += "Plong"
	}

	if rain == "" {
		rain = strconv.Itoa(num)
	}
	return rain
}
