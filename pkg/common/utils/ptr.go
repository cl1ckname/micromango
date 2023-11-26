package utils

func DerefOrDefault[T any](p *T, d T) T {
	if p == nil {
		return d
	}
	return *p
}

func Ptr[T any](v T) *T {
	return &v
}
