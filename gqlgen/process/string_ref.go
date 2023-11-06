package process

func stringRef(s string) *string {
	if s == "" {
		return nil
	} else {
		return &s
	}
}
