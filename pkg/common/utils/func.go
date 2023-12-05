package utils

func Map[M, T any](inp []M, f func(M) T) []T {
	r := make([]T, len(inp))
	for i := 0; i < len(inp); i++ {
		r[i] = f(inp[i])
	}
	return r
}
