package set1

import "fmt"

func FixedXOR(a []byte, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices must be of the same length, a: %v, b: %v", len(a), len(b))
	}
	ans := make([]byte, len(a))
	for i := range len(a) {
		ans[i] = a[i] ^ b[i]
	}
	return ans, nil
}
