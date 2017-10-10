package challenge6

import "testing"

func TestCalculateHammingDistance(t *testing.T) {
	if CalculateHammingDistance([]byte("this is a test"), []byte("wokka wokka!!!")) != 37 {
		t.Fatal("Hamming Distance function failure")
	}
}

func TestProblem6(t *testing.T) {

}
