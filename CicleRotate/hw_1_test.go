package ciclerotate

import "testing"

func TestSolution1(t *testing.T) {
	A := [5]int{3, 8, 9, 7, 6}
	R := [5]int{9, 7, 6, 3, 8}
	K := 3
	B := Solution(A[:], K)
	for i := 0; i < len(A); i++ {
		if R[i] != B[i] {
			t.Fatal("Ошибка на тесте 1")
		}
	}
}

func TestSolution2(t *testing.T) {
	A := [4]int{1, 2, 3, 4}
	R := [4]int{1, 2, 3, 4}
	K := 4
	B := Solution(A[:], K)
	for i := 0; i < len(A); i++ {
		if R[i] != B[i] {
			t.Fatal("Ошибка на тесте 1")
		}
	}
}
