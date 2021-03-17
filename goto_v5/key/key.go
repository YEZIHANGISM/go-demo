package key

// KEYCHAR keychar
var KEYCHAR = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenKey 生成短URL
func GenKey(n int) string {
	if n == 0 {
		return string(KEYCHAR[0])
	}

	l := len(KEYCHAR)
	s := make([]byte, 20)
	i := len(s)
	for n > 0 && i >= 0 {
		i--
		j := n % l
		n = (n - j) / l
		s[i] = KEYCHAR[i]
	}
	return string(s[i:])
}
