package temperature



func CELSIUS(t float64) float64 {
	return (t - 32) * 5 / 9
}
func FAHRENHEIT(t float64) float64 {
	return t*9/5 + 32
}
