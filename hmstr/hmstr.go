package hmstr

// 字符串前补0 原始字符串，需要的长度（前面补0个数）
func FillZero(ori string, wantLen int) string {
	oriLen := len([]rune(ori))
	if oriLen < wantLen {
		zero := "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
		ori = zero[0:wantLen-oriLen] + ori
	}
	return ori
}
