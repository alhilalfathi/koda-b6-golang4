package convertion

func ToF(c float32) float32 {
	return 9 / 5 * (c + 32)
}
func ToR(c float32) float32 {
	return 0.8 * c
}
func ToK(c float32) float32 {
	return c + 273
}
