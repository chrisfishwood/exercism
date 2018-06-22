package raindrops

import (
	"math"
	"strconv"
)

//Convert - converts numbers to raindrops :mindblown:
func Convert(num int) string {
	rain := ""
	rain += AddStringIfFactor(num, 3, "Pling")
	rain += AddStringIfFactor(num, 5, "Plang")
	rain += AddStringIfFactor(num, 7, "Plong")

	if rain == "" {
		rain = strconv.Itoa(num)
	}
	return rain
}

//AddStringIfFactor - returns 'raindrop' if 'num' is evenly divisible by 'divisor'
func AddStringIfFactor(num int, divisor int, raindrop string) string {
	rain := ""
	if Remainder(num, divisor) == 0 {
		rain += raindrop
	}
	return rain
}

//Remainder - returns the remainder of 'num' divided by 'divisor'
func Remainder(num int, divisor int) float64 {
	return math.Remainder(float64(num), float64(divisor))
}
