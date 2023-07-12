package internal

func StringRef(s string) *string {
	if s == "" {
		return nil
	} else {
		return &s
	}
}
