package hamming

import "errors"

func Distance(a, b string) (int, error) {
		if len(a) != len(b) {
		//	error := fmt.Sprintf("Strands must be the same length\n%s\n%s", dnaA, dnaB)
		return 0, errors.New("strands must be the same length")
		}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
