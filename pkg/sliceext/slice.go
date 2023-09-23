package sliceext

func Flatten[T any](lists [][]T) (res []T) {
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}

func Filter[T any](s []T, fn func(T) bool) (res []T) {
	for _, e := range s {
		if fn(e) {
			res = append(res, e)
		}
	}
	return res
}

func FindStr(s []string, target string) string {
	for _, e := range s {
		if e == target {
			return e
		}
	}
	return ""
}

func FindAny(s []any, target any) any {
	for _, e := range s {
		if e == target {
			return &e
		}
	}
	return nil
}
