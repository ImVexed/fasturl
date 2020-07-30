// +build gofuzz

package fasturl

func Fuzz(data []byte) int {
	if _, err := ParseURL(string(data)); err != nil {
		return 0
	}
	return 1
}
