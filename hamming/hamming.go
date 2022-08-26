package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Length of input doesn't match")
	}
	hd := 0
	for i := range a {
		if a[i] != b[i] {
			hd++
		}
	}
	return hd, nil
}
