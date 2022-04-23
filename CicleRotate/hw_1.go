package ciclerotate

func Solution(A []int, K int) []int {
	pos := K
	if len(A) == K {
		return A
	}
	if len(A) > K {
		pos = len(A) % K
	}
	B := make([]int, len(A))
	copy(B, A[pos:])
	R := B[pos+1:]
	copy(R, A[0:pos])
	return B
}
