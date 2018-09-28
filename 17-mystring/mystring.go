package mystring

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	// Excercise
}

func Contains(s, substr string) bool {
	// Excercise
}
